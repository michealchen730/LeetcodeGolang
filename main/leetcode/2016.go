package main

func maximumDifference(nums []int) int {
	s := nums[0]
	res := -1
	for k := 1; k < len(nums); k++ {
		if nums[k]-s > 0 && nums[k]-s > res {
			res = nums[k] - s
		}
		if nums[k] < s {
			s = nums[k]
		}
	}
	return res
}
