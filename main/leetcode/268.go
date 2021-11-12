package main

func missingNumber268(nums []int) int {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] == n || nums[i] == i {
			i++
		} else {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for j := 0; j < n; j++ {
		if nums[j] != j {
			return j
		}
	}
	return n
}
