package main

import (
	"sort"
)

type ints1356 struct {
	arr  []int
	bits []int
}

func (s ints1356) Len() int { return len(s.arr) }
func (s ints1356) Less(i, j int) bool {
	if s.bits[s.arr[i]] < s.bits[s.arr[j]] {
		return true
	} else if s.bits[s.arr[i]] > s.bits[s.arr[j]] {
		return false
	} else {
		return s.arr[i] < s.arr[j]
	}
}
func (s ints1356) Swap(i, j int) { s.arr[i], s.arr[j] = s.arr[j], s.arr[i] }

func sortByBits(arr []int) []int {
	bits := make([]int, 10001)
	for i := 0; i < len(bits); i++ {
		bits[i] = bits[i>>1] + i&1
	}
	sort.Sort(ints1356{arr: arr, bits: bits})
	return arr
}
