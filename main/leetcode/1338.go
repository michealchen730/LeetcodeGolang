package main

import "sort"

func minSetSize(arr []int) int {
	mp := make(map[int]int)
	for _, v := range arr {
		mp[v]++
	}
	var val []int
	for _, v := range mp {
		val = append(val, v)
	}
	sort.Ints(val)
	temp := 0
	for i := len(val) - 1; i >= 0; i-- {
		temp += val[i]
		if temp >= len(arr)/2 {
			return len(val) - i
		}
	}
	return 0
}
