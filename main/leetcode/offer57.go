package main

func twoSum(nums []int, target int) []int {
	mp := make(map[int]bool)
	for _, v := range nums {
		mp[v] = true
	}
	for _, v := range nums {
		if mp[target-v] {
			return []int{v, target - v}
		}
	}
	return []int{0, 0}
}
