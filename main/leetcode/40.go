package main

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	used := make([]int, len(candidates))
	var path []int
	var res [][]int
	getCombinationSum2(0, candidates, used, 0, target, &path, &res)
	return res
}

/**
i:当前第i位
*/
func getCombinationSum2(i int, candidates []int, used []int, temp int, target int, path *[]int, res *[][]int) {
	if i == len(candidates) {
		return
	}
	getCombinationSum2(i+1, candidates, used, temp, target, path, res)
	if i == 0 || i > 0 && ((candidates[i] != candidates[i-1]) || (candidates[i] == candidates[i-1] && used[i-1] == 1)) {
		if temp+candidates[i] > target {
			return
		}
		if temp+candidates[i] == target {
			*path = append(*path, candidates[i])
			cpy := make([]int, len(*path))
			copy(cpy, *path)
			*res = append(*res, cpy)
			temp2 := *path
			temp2 = temp2[:len(temp2)-1]
			*path = temp2
			return
		}
		if temp+candidates[i] < target {
			*path = append(*path, candidates[i])
			used[i] = 1
			getCombinationSum2(i+1, candidates, used, temp+candidates[i], target, path, res)
			used[i] = 0
			temp2 := *path
			temp2 = temp2[:len(temp2)-1]
			*path = temp2
		}
	}
}
