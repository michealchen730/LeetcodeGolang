package main

func search704(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 探讨小于号下的mid情况
	// [1,2] target为2 会返回-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
