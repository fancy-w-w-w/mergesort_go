package redismq

import (
	"fmt"
	"woaini/bytes"
)

func Produce(filename string,queueName string){
	nums,err:=bytes.ReadFile(filename)
	if err!=nil{
		fmt.Println(err)
	}
	BatchPushQueue(queueName,nums)
}

