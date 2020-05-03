package main

import (
	"sort"
)

type Ints1387 [][]int

func (s Ints1387) Len() int { return len(s) }
func (s Ints1387) Less(i, j int) bool {
	if s[i][1] < s[j][1] {
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
func (s Ints1387) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func getKth1387(lo int, hi int, k int) int {
	var res [][]int
	for i := lo; i <= hi; i++ {
		res = append(res, []int{i, getWeight1387(i)})
	}
	sort.Sort(Ints1387(res))
	return res[k-1][0]
}

func getWeight1387(x int) int {
	res := 0
	for x != 1 {
		if x%2 == 0 {
			x = x / 2
		} else {
			x = x*3 + 1
		}
		res++
	}
	return res
}
