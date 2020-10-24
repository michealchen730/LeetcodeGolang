package main

import (
	"sort"
)

func change(amount int, coins []int) int {
	arr := make([]int, amount+1)
	arr[0] = 1
	sort.Ints(coins)

	for _, v := range coins {
		for i := 1; i <= amount; i++ {
			if v <= i {
				arr[i] += arr[i-v]
			}
		}
	}
	return arr[amount]
}
