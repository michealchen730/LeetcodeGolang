package main

import "sort"

func getLastMoment(n int, left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	l, r := 0, 0
	if len(left) > 0 {
		l = left[len(left)-1] - 0
	}
	if len(right) > 0 {
		r = n - right[0]
	}
	return max(l, r)
}
