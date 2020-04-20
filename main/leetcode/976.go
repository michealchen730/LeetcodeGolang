package main

import "sort"

func largestPerimeter(A []int) int {
	res := 0
	sort.Ints(A)
ON:
	for i := len(A) - 1; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
		TH:
			for k := j - 1; k >= 0; k-- {
				if A[k]+A[j] <= A[i] {
					break TH
				} else {
					if A[k]+A[j]+A[i] > res {
						res = A[k] + A[j] + A[i]
						break ON
					}
				}
			}
		}
	}
	return res
}
