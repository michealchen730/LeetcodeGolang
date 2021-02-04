package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	res := 1
	t := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			t++
			res = max(res, t)
		} else {
			t = 1
		}
	}
	return res
}
