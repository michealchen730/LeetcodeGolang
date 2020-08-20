package main

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 != 0 {
			i++
			continue
		}
		if nums[j]%2 == 0 {
			j--
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
