package main

import (
	"container/heap"
	"strconv"
	"strings"
)

type MaxHeap struct {
	heap []int
	size int
}

func (this *MaxHeap) push(num int) {
	this.heap = append(this.heap, num)
	this.size++
	k := this.size - 1
	for k > 0 {
		child := (k - 1) / 2
		if this.heap[child] < this.heap[k] {
			this.heap[child], this.heap[k] = this.heap[k], this.heap[child]
			k = child
		} else {
			break
		}
	}
}

func (this *MaxHeap) pop() int {
	res := this.heap[0]
	this.heap[0] = this.heap[this.size-1]
	this.heap = this.heap[:this.size-1]
	this.size--
	this.adjustMaxHeap(0)
	return res
}

func (this *MaxHeap) adjustMaxHeap(i int) {
	for k := 2*i + 1; k < this.size; k = 2*k + 1 {
		if k+1 < this.size && this.heap[k] < this.heap[k+1] {
			k = k + 1
		}
		if this.heap[k] > this.heap[i] {
			this.heap[k], this.heap[i] = this.heap[i], this.heap[k]
			i = k
		} else {
			break
		}
	}
}

type MinHeap struct {
	heap []int
	size int
}

func (this *MinHeap) push(num int) {
	this.heap = append(this.heap, num)
	this.size++
	k := this.size - 1
	for k > 0 {
		child := (k - 1) / 2
		if this.heap[child] > this.heap[k] {
			this.heap[child], this.heap[k] = this.heap[k], this.heap[child]
			k = child
		} else {
			break
		}
	}
}

func (this *MinHeap) pop() int {
	res := this.heap[0]
	this.heap[0] = this.heap[this.size-1]
	this.heap = this.heap[:this.size-1]
	this.size--
	this.adjustMinHeap(0)
	return res
}

func (this *MinHeap) adjustMinHeap(i int) {
	for k := 2*i + 1; k < this.size; k = 2*k + 1 {
		if k+1 < this.size && this.heap[k] > this.heap[k+1] {
			k = k + 1
		}
		if this.heap[k] < this.heap[i] {
			this.heap[k], this.heap[i] = this.heap[i], this.heap[k]
			i = k
		} else {
			break
		}
	}
}

type S struct {
	name  string
	count int
}

type Strs []S

func (h Strs) Len() int      { return len(h) }
func (h Strs) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Strs) Less(i, j int) bool {
	if h[i].count < h[j].count {
		return true
	} else if h[i].count > h[j].count {
		return false
	} else {
		return strings.Compare(h[i].name, h[j].name) > 0
	}
}

func (h *Strs) Push(x interface{}) {
	*h = append(*h, x.(S))
}
func (h *Strs) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKstrings(strings []string, k int) [][]string {
	// write code here
	mp := make(map[string]int)
	for _, v := range strings {
		mp[v]++
	}
	h := &Strs{}
	heap.Init(h)
	for key, v := range mp {
		heap.Push(h, S{name: key, count: v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([][]string, k)
	for i := k - 1; i >= 0; i-- {
		element := heap.Pop(h)
		res[i] = []string{element.(S).name, strconv.Itoa(element.(S).count)}
	}
	return res
}
