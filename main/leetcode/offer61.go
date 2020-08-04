package main

import "sort"

func isStraight(nums []int) bool {
	sort.Ints(nums)
	mx := nums[4]
	mn := nums[4]
	for i := 3; i >= 0; i-- {
		//遇见大小王直接退出
		if nums[i] == 0 {
			break
		}
		//防止出现重复
		if nums[i] == nums[i+1] {
			return false
		}
		mn = min(nums[i], mn)
	}
	if mx-mn <= 4 {
		return true
	} else {
		return false
	}
}
