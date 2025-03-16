// 用 Fix 函数实现一个简单的堆排序算法
package main

import (
	"container/heap"
	"fmt"
)

// IntHeap　是一个整数堆的实现
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 向堆中添加元素
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop从堆中移除并返回最小元素
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// heapSort 使用堆排序堆数组进行排序
func heapSort(arr []int) []int {
	h := IntHeap(arr)
	heap.Init(&h)

	sorted := make([]int, 0, len(arr))
	for h.Len() > 0 {
		//取出堆顶元素
		sorted = append(sorted, heap.Pop(&h).(int))
	}
	return sorted
}

// 使用Fix函数更新堆元素并重新排序
func updateAndSort(h *IntHeap, index, newValue int) []int {
	//更新指定索引处的元素值
	(*h)[index] = newValue
	//使用Fix函数重新调整堆的顺序
	heap.Fix(h, index)

	sorted := make([]int, 0, h.Len())
	for h.Len() > 0 {
		sorted = append(sorted, heap.Pop(h).(int))
	}
	return sorted
}

func main() {
	arr := []int{3, 2, 1, 5, 6, 4}
	//初始化堆
	h := IntHeap(arr)
	heap.Init(&h)
	fmt.Println("初始堆: ", h)
	//更新索引为2的元素值为7
	indexToUpdate := 2
	newValue := 7
	sorted := updateAndSort(&h, indexToUpdate, newValue)
	fmt.Println("更新元素后排序结果: ", sorted)
}
