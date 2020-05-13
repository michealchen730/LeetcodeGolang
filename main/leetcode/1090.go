package main

import (
	"sort"
)

type Ints1090 [][]int

func (s Ints1090) Len() int { return len(s) }
func (s Ints1090) Less(i, j int) bool {
	if s[i][0] > s[j][0] {
		return true
	} else {
		return false
	}
}
func (s Ints1090) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	res := 0
	var arr [][]int
	for k, v := range values {
		arr = append(arr, []int{v, labels[k]})
	}
	sort.Sort(Ints1090(arr))
	mp := make(map[int]int)
	nums := 0
	for _, v := range arr {
		if mp[v[1]] < use_limit {
			res += v[0]
			mp[v[1]]++
			nums++
		}
		if nums == num_wanted {
			break
		}
	}
	return res
}
