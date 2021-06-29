package mergetask

import (
	"log"
	"woaini/mergesort/algorithm"
)

type SortTask struct {
	sorter  *algorithm.Quick
	retChan chan *MinInt64Slice
}

func NewSortTask(src []int64, retChan chan *MinInt64Slice) *SortTask {

	return &SortTask{
		sorter: algorithm.NewQuick(&MinInt64Slice{
			array: src,
		}),
		retChan: retChan,
	}
}

func (s *SortTask) Run() error {
	s.sorter.Sort()
	// 运行结束返回结果
	s.retChan <- s.sorter.GetSortable().(*MinInt64Slice)
	return nil
}

func (s *SortTask) RunError(err error) {
	log.Println("sort task run error", err)
}

type MergeTask struct {
	slices  [][]int64
	retChan chan []int64
}

func NewMergeTask(slices [][]int64, retChan chan []int64) *MergeTask {
	return &MergeTask{
		slices:  slices,
		retChan: retChan,
	}
}

func (m *MergeTask) Run() error {
	sortedSlices := make([]*algorithm.SortedSlice, 0, len(m.slices))
	for _, s := range m.slices {
		sortedSlices = append(sortedSlices, algorithm.NewSortedSlice(s))
	}	//构造新的sortedSlices,该指针数组里面包含sortedSlice的指针

	//将sortedSlice包装成heap(*HeapMerge)
	merge := algorithm.NewHeapMerge(sortedSlices)


	//merge.Build()
	m.retChan <- merge.Sort()
	return nil
}

func (m *MergeTask) RunError(err error) {
	log.Println("MergeTask run error", err)
}
