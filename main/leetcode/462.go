package main

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for _, v := range nums {
		sum += abs(nums[len(nums)/2] - v)
	}
	return sum
}
