package main

import (
	"sort"
)

//两层循环套双指针
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int

	//枚举所有第一个数
	for first := 0; first < len(nums); first++ {
		//这里做个基本的剪枝
		if nums[first]*4 > target {
			break
		}
		//避免重复枚举
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		//枚举第二个数
		for second := first + 1; second < len(nums); second++ {
			//避免重复枚举
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			//双指针部分
			last := len(nums) - 1

			//这里看起来是枚举，实际上属于左边指针的右移过程
			for third := second + 1; third < len(nums); third++ {
				//还是注意避免重复枚举
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}
				//右指针在左指针之前
				for last > third && nums[third]+nums[first]+nums[second]+nums[last] > target {
					last--
				}
				if third == last {
					break
				}
				if nums[third]+nums[first]+nums[second]+nums[last] == target {
					res = append(res, []int{nums[first], nums[second], nums[third], nums[last]})
				}
			}
		}
	}
	return res
}
