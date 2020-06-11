package main

func shuffle(nums []int, n int) []int {
	i, j := 0, n
	var res []int
	for i < n {
		res = append(res, nums[i])
		res = append(res, nums[j])
		i++
		j++
	}
	return res
}
