package main

func majorityElementO39(nums []int) int {
	count := 0
	res := 0
	for k, v := range nums {
		if count == 0 {
			res = v
			count++
			continue
		}
		if v != nums[k-1] {
			count--
		} else {
			count++
		}
	}
	return res
}
