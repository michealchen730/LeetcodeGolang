package main

func moveZeroes(nums []int) {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt++
			j := i + 1
			for ; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[j], nums[j-cnt] = nums[j-cnt], nums[j]
				} else {
					i = j - 1
					break
				}
			}
			if j == len(nums) {
				break
			}
		}
	}
}
