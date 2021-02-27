package main

import "sort"

//背包问题（类似于完全背包）
func combinationSum4(nums []int, target int) int {
	arr := make([]int, target+1)
	arr[0] = 1
	sort.Ints(nums)
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if i >= v {
				arr[i] += arr[i-v]
			} else {
				break
			}
		}
	}
	return arr[target]
}
