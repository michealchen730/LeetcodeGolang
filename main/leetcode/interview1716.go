package main

func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) >= 2 {
		nums[1] = max(nums[1], nums[0])
	}
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-1], nums[i-2]+nums[i])
	}
	return nums[len(nums)-1]
}
