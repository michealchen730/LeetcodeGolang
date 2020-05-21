package main

import (
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	//第一种思路，模拟
	arr := make([]int, len(deck))
	for k, _ := range arr {
		arr[k] = k
	}
	var stack []int
	for len(arr) > 0 {
		stack = append(stack, arr[0])
		arr = arr[1:]
		if len(arr) > 1 {
			arr = append(arr, arr[0])
			arr = arr[1:]
		}
	}
	sort.Ints(deck)
	res := make([]int, len(stack))
	for k, v := range stack {
		res[v] = deck[k]
	}
	return res
}

//第二种思路，数学，找规律，难度太高放弃
