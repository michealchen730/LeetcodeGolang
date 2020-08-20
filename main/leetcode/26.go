package main

//双指针
func removeDuplicates26(nums []int) int {
	left, right := 0, 0
	for right < len(nums) {
		if right == 0 || nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
