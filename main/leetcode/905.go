package main

//双指针
func sortArrayByParity(nums []int) []int {
	first, last := 0, len(nums)-1
	for first < last {
		if nums[first]%2 != 0 && nums[last]%2 == 0 {
			nums[first], nums[last] = nums[last], nums[first]
			first++
			last--
		}
		if nums[first]%2 == 0 {
			first++
		}
		if nums[last]%2 != 0 {
			last--
		}
	}
	return nums
}
