package main

import "sort"

func sequentialDigits(low int, high int) []int {
	var arr []int
	var dfs func(num int, i int)
	dfs = func(num int, i int) {
		if i > 9 {
			return
		}
		num *= 10
		num += i
		if num > high {
			return
		}
		if num >= low {
			arr = append(arr, num)
		}
		dfs(num, i+1)
	}

	for i := 1; i <= 9; i++ {
		dfs(0, i)
	}
	sort.Ints(arr)
	return arr
}
