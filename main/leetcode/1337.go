package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	var s [][]int
	for i := 0; i < len(mat); i++ {
		t := 0
		for j := 0; j < len(mat[0]); j++ {
			t += mat[i][j]
		}
		s = append(s, []int{i, t})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i][1] < s[j][1] || (s[i][1] == s[j][1] && s[i][0] < s[j][0])
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = s[i][0]
	}
	return res
}
