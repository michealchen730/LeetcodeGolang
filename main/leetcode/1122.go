package main

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	mp := make(map[int]int)
	for _, v := range arr1 {
		mp[v]++
	}
	var res []int

	for _, v := range arr2 {
		for i := 0; i < mp[v]; i++ {
			res = append(res, v)
		}
		delete(mp, v)
	}

	var arr3 []int

	for k, v := range mp {
		for i := 0; i < v; i++ {
			arr3 = append(arr3, k)
		}
	}
	sort.Ints(arr3)

	res = append(res, arr3...)
	return res

}
