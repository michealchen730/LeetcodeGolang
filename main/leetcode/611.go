package main

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j] > nums[k] {
					res++
				} else {
					break
				}
			}
		}
	}
	return res
}
