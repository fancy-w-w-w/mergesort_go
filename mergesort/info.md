# Merge Sort

## 问题描述
Go 语言实现一个16M的整数(int64)多路归并的数组排序

## 思路
将待排序数组分成多个组，利用多个goroutine实现各个组的并行排序；然后通过Heap(最小堆)进行多路归并排序；

## 实现
实现一个协程池实现任务的并行处理，将待排序切片分组并封装成SortTask放入协程池
运行，待全部执行完成后ConcurrentSorter收集排序结果，并封装成MergeTask放入协程池中进行合并。

+ 协程池pool.go

    - 配置最大协程数量
    - 按需创建协程
    - 空闲超时则回收协程

+ 合并有序切片algorithm.heapmerge.go
  通过堆实现多路的有序切片的合并，额外申请一倍的内存用于存放合并结果
  
+ 对切片进行快速排序algorithm.quicksort.go
  通过快排得到有序的切片，使用内置sort（）函数节省时间
  
+ 封装slice待归并切片mergetask.slice.go
  
+ service层 mergetask.singletask.go
  实现pool.go中Run()接口，实现快速排序和堆排序的任务处理
  
+ concurrenttask.go  并发多路排序
输入：n路待合并的有序slice
输出：有序slice

堆node定义为一个SortedSlice，实现了hasNext函数，用于迭代到当前slice的下一个元素；
1. 构建一个n个元素的最小堆
2. 从每路slice中取首个元素组成数组，调整堆；每次从堆顶，取一个元素，放入合并后的slice中
    + 如果hasNext=true，执行当前node的Next()，重新调整当前的堆
    + 如果hasNext=false, 当前slice已经空了，因此剔除堆顶, 然后需要重建堆，原因是堆中的父子关系已经破坏。




