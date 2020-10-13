package main

import "sort"

func maxCoins(piles []int) int {
	sort.Ints(piles)
	all, cnt := len(piles)/3, 0
	sum, k := 0, len(piles)-2
	for cnt < all {
		sum += piles[k]
		k -= 2
		cnt++
	}
	return sum
}
