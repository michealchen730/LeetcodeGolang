package main

func countKDifference(nums []int, k int) int {
	sum := 0
	mp := make(map[int]int)
	for _, v := range nums {
		sum += mp[v]
		mp[v-k]++
		mp[v+k]++
	}
	return sum
}
