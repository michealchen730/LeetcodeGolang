package main

import (
	"sort"
)

//第一种解法
//search方法，如果一个数组可以被划分为相同子集和，那么
//1、对数组排序
//2、从后往前，对每个数进行处理，放到盒子里
//3、这里是穷举了所有的可能性，并不是我原先理解的，瓶子里放石头放沙子放水
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	sort.Ints(nums)
	row := len(nums) - 1
	if nums[row] > target {
		return false
	}
	for row >= 0 && nums[row] == target {
		row--
		k--
	}
	arr := make([]int, k)
	return search698(arr, nums, row, target)
}

func search698(groups, nums []int, row, target int) bool {
	if row < 0 {
		return true
	}
	v := nums[row]
	row--
	for i := 0; i < len(groups); i++ {
		if groups[i]+v <= target {
			groups[i] += v
			if search698(groups, nums, row, target) {
				return true
			}
			groups[i] -= v
		}
		if groups[i] == 0 {
			break
		}
	}
	return false
}
