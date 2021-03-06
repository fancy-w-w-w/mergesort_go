package taskpool

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	QSize   int
	Workers int
	MaxIdle time.Duration
}

type Task interface {
	Run() error
	// 当task运行报错时，调用RunError；等价于Run返回一个 chan error 异步监听；
	RunError(err error)
}

type Pool struct {
	Config
	waitingChan chan Task
	workersChan chan struct{} // worker Chan;如果没有达到上限；
	wg          sync.WaitGroup
	close       int32 //atomic
}

func NewPool(cfg *Config) *Pool {

	return &Pool{
		Config:      *cfg,
		waitingChan: make(chan Task, cfg.QSize),
		workersChan: make(chan struct{}, cfg.Workers),
		wg:          sync.WaitGroup{},
		close:       0,
	}
}

func (p *Pool) Put(t Task) *Pool {
	select {
	// 按需创建：如果worker数量未到上限,就创建一个worker；
	// 优先放入workersChan
	case p.workersChan <- struct{}{}:
		p.wg.Add(1)
		go p.work(t)
	default:
		select {
		case p.workersChan <- struct{}{}:
			p.wg.Add(1)
			go p.work(t)
		case p.waitingChan <- t:
		}
	}
	return p
}

func (p *Pool) work(t Task) {
	// 最长等待时间

	timer := time.NewTimer(p.MaxIdle)

	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error: // 运行时错误
				fmt.Println("goroutine pool:", err)
			default: // 非运行时错误
				fmt.Println("error:", err)
			}
		}

	}()
	defer func() {
		timer.Stop()
		p.wg.Done()
		// worker退出之后，从worker chan读；当有新的任务的时候，可以直接启动新的worker
		<-p.workersChan
		//p.exit <- w.name //当前的worker退出，从pool中删除这个worker
	}()
	for {
		if err := t.Run(); err != nil {
			t.RunError(err)
		}
		select {
		// 超时回收：超时仍未收到新的任务，退出当前的任务
		case <-timer.C:
			return
		case newTask, ok := <-p.waitingChan:
			if !ok {
				return
			}
			if newTask == nil {
				return
			}

			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(p.MaxIdle)
			//继续执行新的task
			t = newTask
		}
	}
}

func (p *Pool) Close(grace bool) {
	if ok := atomic.CompareAndSwapInt32(&p.close, 0, 1); ok {
		close(p.waitingChan)
		close(p.workersChan)
		if grace {
			p.wg.Wait()
		}
	}
}
