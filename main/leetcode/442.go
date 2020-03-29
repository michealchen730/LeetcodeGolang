package main

func findDuplicates(nums []int) []int {
	var res []int
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] < 0 {
			res = append(res, v)
		} else {
			nums[v-1] = -nums[v-1]
		}
	}
	return res
}
