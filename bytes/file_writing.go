package bytes

import (
	"fmt"
	"sync"
)

func Write(filename string){


	data := make(chan int64)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go ProduceFile(data, &wg)
	}
	go ConsumeFile(data,done,filename)
	go func() {
		wg.Wait()
		close(data)  //所有协程完成随机数写入然后关闭data
		//当执行了 close(data) ，那么就不能向管道发送信息了，但是在关闭之前留在管道里面的消息还可以被消费。
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}

}