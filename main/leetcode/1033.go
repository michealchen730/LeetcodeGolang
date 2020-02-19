package main

import "sort"

func numMovesStones(a int, b int, c int) []int {
	arr := []int{a, b, c}
	var res []int
	sort.Ints(arr)
	minstps := 2
	if arr[1]-arr[0] == 1 && arr[2]-arr[1] == 1 {
		minstps = 0
	} else {
		for i := 1; i < len(arr); i++ {
			if arr[i]-arr[i-1] <= 2 {
				minstps = 1
				break
			}
		}
	}
	res = append(res, minstps)
	res = append(res, arr[2]-arr[0]-2)
	return res
}
