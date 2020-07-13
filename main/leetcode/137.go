package main

func singleNumber137(nums []int) int {
	mp := make(map[int]int)
	for _, v := range nums {
		if mp[v] < 2 {
			mp[v]++
		} else {
			delete(mp, v)
		}
	}
	for k, _ := range mp {
		return k
	}
	return -1
}
