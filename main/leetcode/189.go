package main

func rotate189(nums []int, k int) {
	if k == len(nums) {
		return
	}
	k %= len(nums)
	reverseArr(nums)
	reverseArr(nums[0:k])
	reverseArr(nums[k:])
}
