package mergesort

import "woaini/mergesort/mergetask"

func MergeSort(src []int64) {

	if len(src) > mergetask.SORTING_ARRAY_THRESHOLD {
		sorter := mergetask.NewConcurrent(src)
		sorter.Run()
	} else {
		sorter := mergetask.NewSingle(src)
		sorter.Run()
	}
}
