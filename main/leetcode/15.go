package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	if len(nums) < 3 || nums[0] > 0 || nums[len(nums)-1] < 0 {
		return res
	}
	for i := 0; i < len(nums)-2; i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] || nums[i]+nums[len(nums)-1]+nums[len(nums)-2] < 0 {
			continue
		}
		if nums[i] > 0 {
			break
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[j] > 0 && nums[i]+nums[j] >= 0 {
				break
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					break
				}
				if nums[i]+nums[j]+nums[k] > 0 {
					break
				}
			}
		}
	}
	return res
}
