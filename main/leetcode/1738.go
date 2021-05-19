package main

import (
	"container/heap"
)

func kthLargestValue(matrix [][]int, k int) int {
	var h hpe
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i != 0 || j != 0 {
				tmp := 0
				if j > 0 {
					tmp ^= matrix[i][j-1]
				}
				if i > 0 {
					tmp ^= matrix[i-1][j]
				}
				if i > 0 && j > 0 {
					tmp ^= matrix[i-1][j-1]
				}
				matrix[i][j] ^= tmp
			}
			heap.Push(&h, matrix[i][j])
			if len(h) == k {
				heap.Init(&h)
			}
			if len(h) > k {
				heap.Pop(&h)
			}
		}
	}
	return heap.Pop(&h).(int)
}

type hpe []int

func (h hpe) Len() int {
	return len(h)
}

func (h hpe) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h hpe) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hpe) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *hpe) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
