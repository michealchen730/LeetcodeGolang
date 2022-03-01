package main

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	sort.Ints(nums)
	res := math.MaxInt32
	for t := k - 1; t < len(nums); t++ {
		if nums[t]-nums[t-k+1] < res {
			res = nums[t] - nums[t-k+1]
		}
	}
	return res
}
