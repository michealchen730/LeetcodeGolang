package main

func findRepeatNumber(nums []int) int {
	mp := make(map[int]int)
	for _, v := range nums {
		if mp[v] > 0 {
			return v
		} else {
			mp[v]++
		}
	}
	return -1
}
