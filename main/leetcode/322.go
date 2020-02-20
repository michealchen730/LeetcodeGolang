package main

import (
	"sort"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	sort.Ints(coins)
	arr := make([]int, amount+1)
	for i := 0; i < len(coins); i++ {
		if coins[i] <= amount {
			arr[coins[i]] = 1
		}
	}
	return getCoinChange(coins, amount, arr)
}

func getCoinChange(coins []int, amount int, arr []int) int {
	if arr[amount] != 0 {
		return arr[amount]
	}
	if amount < coins[0] {
		arr[amount] = -1
		return -1
	}
	i := 0
	res := amount + 1
	for i < len(coins) && coins[i] < amount {
		temp := getCoinChange(coins, amount-coins[i], arr)
		if temp != -1 && temp < res {
			res = temp
		}
		i++
	}
	if res == amount+1 {
		arr[amount] = -1
	} else {
		arr[amount] = res + 1
	}
	return arr[amount]
}
