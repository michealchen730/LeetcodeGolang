package main

import "sort"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	var res [][]int
	mp := make(map[int][][]int)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			t := abs(i-r0) + abs(j-c0)
			mp[t] = append(mp[t], []int{i, j})
		}
	}
	maxd := []int{abs(0-r0) + abs(0-c0), abs(R-r0) + abs(0-c0), abs(0-r0) + abs(C-c0), abs(R-r0) + abs(C-c0)}
	sort.Ints(maxd)
	for i := 0; i <= maxd[3]; i++ {
		for j := 0; j < len(mp[i]); j++ {
			res = append(res, mp[i][j])
		}
	}
	return res
}
