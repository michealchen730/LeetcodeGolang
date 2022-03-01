package main

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var arr []int
	if nums[1] < nums[0] {
		arr = []int{1, 0}
	} else {
		arr = []int{0, 1}
	}
	for k := 2; k < len(nums); k++ {
		if nums[k] > nums[arr[1]] {
			arr[0], arr[1] = arr[1], k
			continue
		}
		if nums[k] > nums[arr[0]] {
			arr[0] = k
		}
	}
	if nums[arr[1]]-nums[arr[0]] >= nums[arr[0]] {
		return arr[1]
	}
	return -1
}
