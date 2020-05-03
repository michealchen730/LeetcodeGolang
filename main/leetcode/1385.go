package main

import "sort"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	sort.Ints(arr1)
	sort.Ints(arr2)
	for _, v1 := range arr1 {
		flag := 0
		for _, v2 := range arr2 {
			if abs(v2-v1) <= d {
				flag = 1
				break
			}
		}
		if flag == 0 {
			res++
		}
	}
	return res
}
