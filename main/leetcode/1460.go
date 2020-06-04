package main

import "sort"

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for k, v := range target {
		if v != arr[k] {
			return false
		}
	}
	return true
}
