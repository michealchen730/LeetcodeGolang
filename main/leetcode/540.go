package main

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		width := mid - low
		if nums[mid] == nums[mid+1] {
			//如果两侧的宽度为奇数
			if width%2 == 0 {
				low = mid
			} else {
				high = mid - 1
			}
			continue
		}
		if nums[mid] == nums[mid-1] {
			if width%2 == 0 {
				high = mid
			} else {
				low = mid + 1
			}
			continue
		}
		return nums[mid]
	}
	return nums[low]
}
