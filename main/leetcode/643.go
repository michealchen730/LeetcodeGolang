package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	mn := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			nums[i] += nums[i-1]
		}
		if i == k-1 {
			mn = nums[i]
		}
		if i >= k {
			mn = max(mn, nums[i]-nums[i-k])
		}
	}
	return float64(mn) / float64(k)
}
