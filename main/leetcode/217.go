package main

func containsDuplicate(nums []int) bool {
	mp := make(map[int]int)

	for _, v := range nums {
		mp[v]++
		if mp[v] >= 2 {
			return true
		}
	}
	return false
}
