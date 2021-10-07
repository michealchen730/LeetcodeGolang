package main

import (
	"container/heap"
	"sort"
)

//根据贪心的原则选择

type PC struct {
	profits int
	capital int
}

type hpc []PC

func (h hpc) Less(i, j int) bool {
	return h[i].profits > h[j].profits
}

func (h hpc) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hpc) Len() int {
	return len(h)
}

func (h *hpc) Push(x interface{}) {
	*h = append(*h, x.(PC))
}

func (h *hpc) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	sourcePC := make([]PC, len(profits))
	for idx, v := range profits {
		sourcePC[idx].profits = v
		sourcePC[idx].capital = capital[idx]
	}
	sort.Slice(sourcePC, func(i, j int) bool {
		return sourcePC[i].capital < sourcePC[j].capital
	})
	h := new(hpc)
	heap.Init(h)
	for idx := 0; idx < len(sourcePC); idx++ {
		if sourcePC[idx].capital <= w {
			heap.Push(h, sourcePC[idx])
		} else {
			if len(*h) == 0 {
				break
			}
			t := heap.Pop(h)
			w += t.(PC).profits
			k--
			idx--
			if k == 0 {
				break
			}
		}
	}
	for len(*h) > 0 && k > 0 {
		t := heap.Pop(h)
		w += t.(PC).profits
		k--
	}
	return w
}
