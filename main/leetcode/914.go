package main

import "sort"

//空间复杂度（比map解法要好）
func hasGroupsSizeX(deck []int) bool {
	sort.Ints(deck)
	arr := make([]int, deck[len(deck)-1]-deck[0]+1)
	for _, v := range deck {
		arr[v-deck[0]]++
	}
	sort.Ints(arr)
	gcd := arr[len(arr)-1]
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			break
		}
		temp := getGreatestCommonDivisor(gcd, arr[i])
		if temp < gcd {
			gcd = temp
		}
		if gcd == 1 {
			return false
		}
	}
	return true
}

//优化一
func hasGroupsSizeX2(deck []int) bool {
	mp := make(map[int]int)
	for _, v := range deck {
		mp[v]++
	}
	gcd := mp[deck[0]]
	for _, v := range mp {
		temp := getGreatestCommonDivisor(gcd, v)
		if temp < gcd {
			gcd = temp
		}
		if gcd == 1 {
			return false
		}
	}
	return true
}
