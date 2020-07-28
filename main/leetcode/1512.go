package main

func numIdenticalPairs(nums []int) int {
	res := 0
	mp := make(map[int]int)
	for _, v := range nums {
		res += mp[v]
		mp[v]++
	}
	return res
}
