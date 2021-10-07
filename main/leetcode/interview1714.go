package main

import "container/heap"

//func smallestK(arr []int, k int) []int {
//	var res []int
//	if k == 0 {
//		return res
//	}
//	if k >= len(arr) {
//		return arr
//	}
//	for i := 0; i < k; i++ {
//		res = append(res, arr[i])
//	}
//	buildMaxHeap(res)
//	for i := k; i < len(arr); i++ {
//		if arr[i] < res[0] {
//			res[0] = arr[i]
//			adjustHeap(res, 0, k)
//		}
//	}
//	return res
//}

type H []int

func (h H) Len() int {
	return len(h)
}

func (h H) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h H) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *H) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *H) Pop() interface{} {
	old := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return old
}

func smallestK(arr []int, k int) []int {
	h := H(arr[:k])
	heap.Init(&h)
	for i := k; i < len(arr); i++ {
		heap.Push(&h, arr[i])
		heap.Pop(&h)
	}
	res := make([]int, k)
	for i, _ := range res {
		res[i] = heap.Pop(&h).(int)
	}
	return res
}
