package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	res := make([]int, len(nums))
	res[0] = nums[0]
	if nums[0] > nums[1] {
		res[1] = nums[0]
	} else {
		res[1] = nums[1]
	}
	i := 2
	for i < len(res) {
		if nums[i]+res[i-2] > res[i-1] {
			res[i] = nums[i] + res[i-2]
		} else {
			res[i] = res[i-1]
		}
		i++
	}
	return res[i-1]
}
