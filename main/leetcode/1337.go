package main

import "sort"

type Ints1337 [][]int

func (s Ints1337) Len() int { return len(s) }
func (s Ints1337) Less(i, j int) bool {
	if s[i][1] < s[j][1] {
		return true
	} else if s[i][1] > s[j][1] {
		return false
	} else {
		return i < j
	}
}
func (s Ints1337) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func kWeakestRows(mat [][]int, k int) []int {
	var s [][]int
	for i := 0; i < len(mat); i++ {
		t := 0
		for j := 0; j < len(mat[0]); j++ {
			t += mat[i][j]
		}
		s = append(s, []int{i, t})
	}
	sort.Sort(Ints1337(s))
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k] = s[i][0]
	}
	return res
}
