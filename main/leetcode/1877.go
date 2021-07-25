package main

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	low, high := 0, len(nums)-1
	res := nums[low] + nums[high]
	for low < high {
		res = max(res, nums[low]+nums[high])
		low++
		high--
	}
	return res
}
