package main

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	if nums[low] < nums[high] {
		return nums[low]
	}
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] >= nums[0] {
			if nums[mid] > nums[mid+1] {
				return nums[mid+1]
			}
			low = mid + 1
			continue
		}
		if nums[mid] < nums[0] {
			if nums[mid] < nums[mid-1] {
				return nums[mid]
			}
			high = mid - 1
		}
	}
	return nums[low]
}
