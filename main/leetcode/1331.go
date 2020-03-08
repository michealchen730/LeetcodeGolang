package main

import (
	"sort"
)

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	if len(arr) == 1 {
		return []int{1}
	}
	arr2 := make([]int, len(arr))
	copy(arr2, arr)
	sort.Ints(arr2)
	mp := make(map[int]int)
	temp := 1
	for i := 1; i < len(arr2); i++ {
		mp[arr2[i-1]] = temp
		if arr2[i] != arr2[i-1] {
			temp++
		}
	}
	mp[arr2[len(arr)-1]] = temp
	for i := 0; i < len(arr); i++ {
		arr2[i] = mp[arr[i]]
	}
	return arr2
}
