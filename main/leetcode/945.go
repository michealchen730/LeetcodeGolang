package main

import "sort"

func minIncrementForUnique(A []int) int {
	res := 0
	sort.Ints(A)
	if len(A) <= 1 {
		return res
	}
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			A[i]++
			res++
		} else {
			if A[i] < A[i-1] {
				res += A[i-1] + 1 - A[i]
				A[i] = A[i-1] + 1
			}
		}
	}
	return res
}
