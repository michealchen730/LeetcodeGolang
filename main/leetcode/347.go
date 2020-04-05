package main

import (
	"sort"
)

type Ints [][]int

func (s Ints) Len() int { return len(s) }
func (s Ints) Less(i, j int) bool {
	if s[i][1] > s[j][1] {
		return true
	} else if s[i][1] == s[j][1] {
		if s[i][0] < s[j][0] {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func (s Ints) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	var arr [][]int
	for k, v := range mp {
		arr = append(arr, []int{k, v})
	}
	sort.Sort(Ints(arr))
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, arr[i][0])
	}
	return res
}
