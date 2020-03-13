package main

func majorityElement(nums []int) int {
	temp, count := -1, 0
	for i := 0; i < len(nums); i++ {
		if temp == -1 {
			temp = i
		}
		if nums[i] == nums[temp] {
			count++
		} else {
			count--
		}
		if count > len(nums)/2 {
			return count
		}
		if count == 0 {
			temp = -1
		}
	}
	return nums[temp]
}
