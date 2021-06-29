package redismq

import (
	"fmt"
	"sync"
)

func Consume(queueName string,wg *sync.WaitGroup) int64{
	popnum,err:=PopQueue(queueName,1000)
	if err!=nil{
		fmt.Println("pop error")
	}
	wg.Done()
	return popnum
}

func ConsumeSome(queueName string) []int64{
	wg:=sync.WaitGroup{}
	nums:=make([]int64,0,100000)

	for i:=0;i<100000;i++{
			wg.Add(1)
			j:=Consume(queueName,&wg)
			nums=append(nums,j)
	}
	go func(){
		wg.Wait()
	}()
	return nums
}

func ProduceAndConsume(fileName string,queueName string) []int64{

	Produce(fileName,queueName)
	res:=ConsumeSome(queueName)
	return res
}
