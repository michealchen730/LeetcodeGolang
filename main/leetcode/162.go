package main

func findPeakElement(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if (mid == 0 || nums[mid] > nums[mid-1]) && (mid == len(nums)-1 || nums[mid] > nums[mid+1]) {
			return mid
		}
		if nums[mid] < nums[mid+1] {
			low = mid + 1
			continue
		}
		if nums[mid] < nums[mid-1] {
			high = mid - 1
		}
	}
	return low
}
