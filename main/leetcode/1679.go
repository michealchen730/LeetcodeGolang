package main

func maxOperations(nums []int, k int) int {
	res := 0
	mp := make(map[int]int)
	for _, v := range nums {
		if mp[k-v] == 0 {
			mp[v]++
		} else {
			mp[k-v]--
			res++
		}
	}
	return res
}
