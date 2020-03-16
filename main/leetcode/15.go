package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if j != i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k != len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			temp := nums[i] + nums[j] + nums[k]
			if temp > 0 {
				k--
			}
			if temp < 0 {
				j++
			}
			if temp == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
			}
		}
	}
	return res
}
