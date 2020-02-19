package main

import (
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	sort.Ints(arr)
	m := make(map[int]int)
	temp := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			if _, ok := m[temp]; ok {
				return false
			}
			m[temp] = 1
			temp = 1
		} else {
			temp++
		}
	}
	if _, ok := m[temp]; ok {
		return false
	} else {
		return true
	}
}
