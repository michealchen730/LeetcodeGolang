package main

func findDisappearedNumbers(nums []int) []int {
	for k, v := range nums {
		if v == k+1 {
			continue
		}
		for nums[k] != k+1 && nums[k] != nums[nums[k]-1] {
			nums[k], nums[nums[k]-1] = nums[nums[k]-1], nums[k]
		}
	}
	var res []int
	for k, v := range nums {
		if k+1 != v {
			res = append(res, k+1)
		}
	}
	return res
}
