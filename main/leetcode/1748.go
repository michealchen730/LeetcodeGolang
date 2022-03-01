package main

func sumOfUnique(nums []int) int {
	mp := make(map[int]int)
	sum := 0
	for _, v := range nums {
		if mp[v] == 0 {
			sum += v
		}
		if mp[v] == 1 {
			sum -= v
		}
		mp[v]++
	}
	return sum
}
