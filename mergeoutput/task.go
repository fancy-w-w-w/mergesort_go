package mergeoutput

import (
	"strconv"
	"woaini/mergesort"
	"woaini/mergesort/algorithm"
	"woaini/redismq"
)

func NewSlice(folder string,fileamount int) []*algorithm.SortedSlice{
	var sources []*algorithm.SortedSlice


	for i:=1;i<=fileamount;i++{
		filename:=folder+strconv.Itoa(i)+".txt"
		queuename:="queue"+strconv.Itoa(i)
		num:=redismq.ProduceAndConsume(filename,queuename)
		mergesort.MergeSort(num)
		sources=append(sources,algorithm.NewSortedSlice(num))
	}
	return sources
}

