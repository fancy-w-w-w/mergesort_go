package redismq

import (
	"fmt"
	"testing"
)

func TestProduceAndConsume(t *testing.T) {
	res:=ProduceAndConsume("../bytes/byte1.txt","queue1")

	fmt.Println(res[99999])
}
