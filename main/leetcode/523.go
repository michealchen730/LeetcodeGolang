package main

//给定一个包含非负数的数组和一个目标整数k，编写一个函数来判断该数组是否含有连续的子数组，其大小至少为2，且总和为k的倍数，即总和为 n*k，其中n也是一个整数。

//该连续子数组的长度至少为2
//该连续子数组的和为k的倍数，[0,0]也可以

func checkSubarraySum(nums []int, k int) bool {
	mp := make(map[int]int)
	mp[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if k != 0 {
			sum = sum % k
		}
		if v, ok := mp[sum]; ok {
			if i-v > 1 {
				return true
			}
		} else {
			mp[sum] = i
		}
	}
	return false
}
