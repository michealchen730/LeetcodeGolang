package main

import "sort"

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	m := arr[(len(arr)-1)/2]
	var res []int
	i, j := 0, len(arr)-1
	for len(res) < k {
		if abs(arr[i]-m) > abs(arr[j]-m) {
			res = append(res, arr[i])
			i++
		} else {
			res = append(res, arr[j])
			j--
		}
	}
	return res
}
