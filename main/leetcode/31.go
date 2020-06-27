package main

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := len(nums) - 1

	for i-1 >= 0 && nums[i-1] >= nums[i] {
		i--
	}
	if i == 0 {
		reverseArr(nums)
	} else {
		for temp := len(nums) - 1; temp >= 0; temp-- {
			if nums[temp] > nums[i-1] {
				nums[temp], nums[i-1] = nums[i-1], nums[temp]
				reverseArr(nums[i:len(nums)])
				break
			}
		}
	}
}
