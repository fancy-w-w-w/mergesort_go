package mergesort

import (
	"testing"
	"woaini/bytes"
)

func TestMergeSort(t *testing.T) {
	nums,err:=bytes.ReadFile("../bytes/byte1.txt")
	if err!=nil{
		panic(err)
	}
	MergeSort(nums)
	t.Parallel()
	t.Log(nums[0:100])
}
