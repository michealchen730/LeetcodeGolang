package main

import (
	"math"
	"sort"
)

type Ints406 [][]int

func (s Ints406) Len() int { return len(s) }
func (s Ints406) Less(i, j int) bool {
	if s[i][0] < s[j][0] {
		return true
	} else {
		if s[i][0] == s[j][0] {
			return s[i][1] < s[j][1]
		}
		return false
	}
}
func (s Ints406) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(Ints406(people))
	res := make([][]int, len(people))
	used := make([]bool, len(people))
	for k, _ := range res {
		res[k] = make([]int, 2)
		res[k][0] = math.MinInt32
	}
	for _, v := range people {
		a := 0
		for i := 0; i < len(used); i++ {
			if res[i][0] >= v[0] {
				a++
			}
			if !used[i] {
				a++
			}
			if v[1]+1 == a {
				used[i] = true
				res[i] = v
				break
			}
		}
	}
	return res
}
