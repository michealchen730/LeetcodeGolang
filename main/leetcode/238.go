package main

func productExceptSelf(nums []int) []int {
	R := make([]int, len(nums))
	R[len(R)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	temp := nums[0]
	nums[0] = 1
	for i := 1; i < len(nums); i++ {
		nums[i], temp = nums[i-1]*temp, nums[i]
		//这两个式子没有先后顺序，可以理解为同时发生的
		//如果存在先后顺序，那么temp会变成第一步的nums[i]，但实际上没有
		R[i] *= nums[i]
	}
	return R
}
