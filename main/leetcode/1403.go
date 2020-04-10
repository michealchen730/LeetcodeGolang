package main

import "sort"

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	amount := 0
	for _, v := range nums {
		amount += v
	}
	var res []int
	var amount2 int
	for i := len(nums) - 1; i >= 0; i-- {
		res = append(res, nums[i])
		amount2 += nums[i]
		amount -= nums[i]
		if amount2 > amount {
			break
		}
	}
	return res
}
