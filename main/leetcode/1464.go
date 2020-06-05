package main

func maxProduct1464(nums []int) int {
	res := (nums[0] - 1) * (nums[1] - 1)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			res = max((nums[i]-1)*(nums[j]-1), res)
		}
	}
	return res
}
