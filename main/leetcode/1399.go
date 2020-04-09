package main

import "sort"

func countLargestGroup(n int) int {
	res := make([]int, 37)
	for i := 1; i <= n; i++ {
		res[getSum(i)]++
	}
	sort.Ints(res)
	temp := res[len(res)-1]
	m := 0
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == temp {
			m++
		}
	}
	return m
}

func getSum(n int) int {
	res := 0
	for n != 0 {
		res += n % 10
		n = n / 10
	}
	return res
}
