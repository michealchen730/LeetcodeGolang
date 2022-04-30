package main

import "math"

//看了题解才会的
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for k := 1; k < len(nums); k++ {
		if nums[k] > second {
			return true
		}
		if nums[k] > first {
			second = nums[k]
		}
		if nums[k] < first {
			first = nums[k]
		}
	}
	return false
}
