package bytes

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func ProduceFile(data chan int64, wg *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int63n(999999)+0
	data <- n
	wg.Done()
}


//consume函数创建了一个叫filename的文件
//然后，它从data信道读取随机数并写入文件。一旦读写完所有的随机数，它就会向done信道写入true，通知它已经完成了任务。
func ConsumeFile(data chan int64, done chan bool,filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}


