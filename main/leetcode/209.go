package main

func minSubArrayLen(s int, nums []int) int {
	i, j := 0, 0
	sum := 0
	mx := len(nums) + 1
	for j < len(nums) {
		sum += nums[j]
		for sum >= s {
			mx = min(mx, j-i+1)
			if mx == 1 {
				return mx
			}
			sum -= nums[i]
			i++
		}
		j++
	}
	if mx == len(nums)+1 {
		return 0
	}
	return mx
}
