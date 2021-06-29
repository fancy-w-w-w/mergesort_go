package main

import (
	"fmt"
	"woaini/mergeoutput"
)

func main(){
	//生成新的随机数文件（bytes）

	/*
	var filename string
	filename="../bytes/byte1"
	bytes.Write(filename)
	*/

	fmt.Println("开始执行程序，请稍等，大约1min")

	//函数耗时较长,输出到bytes/merge.txt
	mergeoutput.TaskRun("./bytes/byte","./bytes/merge.txt")


}
