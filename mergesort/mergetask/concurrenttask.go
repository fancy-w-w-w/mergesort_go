package mergetask

import (
	"runtime"
	"sync"
	"time"
	"woaini/mergesort/algorithm"
	"woaini/mergesort/taskpool"
)

const SORTING_ARRAY_THRESHOLD = 1 << 4

type Sorter interface {
	Run()
}

type SingleSorter struct {
	sortingArray []int64
}

func (s *SingleSorter) Run() {
	h := algorithm.NewQuick(&MinInt64Slice{s.sortingArray})
	h.Sort()
}

func NewSingle(src []int64) Sorter {
	return &SingleSorter{sortingArray: src}
}

type ConcurrentSorter struct {
	sortedChan   chan *MinInt64Slice
	sortingArray []int64
	pool         *taskpool.Pool
	taskNum      int
}

func NewConcurrent(src []int64) Sorter {
	// 拆分成子任务并行完成
	taskNum := runtime.NumCPU()
	return &ConcurrentSorter{
		sortedChan: make(chan *MinInt64Slice, 1),
		//mergeRetChan: make(chan []int64),
		sortingArray: src,
		pool: taskpool.NewPool(&taskpool.Config{
			QSize:   1,
			Workers: runtime.NumCPU(),
			MaxIdle: time.Second * 10,
		}),
		taskNum: taskNum,
	}
}

func (m *ConcurrentSorter) sort() {
	start := 0
	step := len(m.sortingArray) / m.taskNum
	// 不能整除，则最后一个task多处理一些
	count := 1
	for ; start < len(m.sortingArray); {
		end := (start + step) % len(m.sortingArray)
		// 最后一个任务
		if m.taskNum == count {
			end = len(m.sortingArray)
		}
		t := NewSortTask(m.sortingArray[start:end], m.sortedChan)
		start = end
		m.pool.Put(t)
		count++
	}
}

func (m *ConcurrentSorter) merge(mergedChan chan []int64) {
	sortedSlices := make([][]int64, 0, m.taskNum)
	sortedLen := 0
loop:
	for {
		select {
		case s := <-m.sortedChan:
			sortedLen += s.Len()
			sortedSlices = append(sortedSlices, s.GetSlice())
			// sort 阶段完成
			if sortedLen == len(m.sortingArray) {
				// 这里确保所有的task都已经退出，不然可能导致死锁
				// 死锁产生的场景，SortTask.Run()->SortedChan,如果该routine退出<-m.sortedChan;
				// 那么SortedTask无法退出；当前task m.pool.Put(MergeTask)就会阻塞
				break loop
			}
		}
	}
	mergeTask := NewMergeTask(sortedSlices, mergedChan)
	// 为避免死锁,另外启动一个协程写入
	m.pool.Put(mergeTask)
}

func (m *ConcurrentSorter) Run() {
	mergedChan := make(chan []int64, 1)
	defer func() {
		m.pool.Close(true)
		close(m.sortedChan)
		close(mergedChan)
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		m.sort()    //开启一个协程执行sort()，将sort()加入任务队列，使用线程池等待处理
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		m.merge(mergedChan)	//开启一个协程执行merge()，将merge()加入任务队列，使用线程池等待处理
		wg.Done()
	}()
	wg.Wait()

	resultSlice := <-mergedChan
	copy(m.sortingArray, resultSlice)
}
