package mergeoutput

import (
	"fmt"
	"os"
	"woaini/mergesort/algorithm"
)

func TaskRun(filepath string,outputpath string){
	var sources []*algorithm.SortedSlice
	var heap *algorithm.HeapMerge


	sources=NewSlice(filepath,10)
	heap=algorithm.NewHeapMerge(sources)
	go heap.Print()
	out:=heap.Sort()  //归并
	f, err := os.Create(outputpath)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _,d := range out {
			_, err = fmt.Fprintln(f, d)
			if err != nil {
				fmt.Println(err)
				f.Close()
		}
	}


}
