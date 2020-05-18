package main

import "sort"

func minFallingPathSum931(A [][]int) int {
	for i := 1; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			temp := A[i][j]
			A[i][j] += A[i-1][j]
			for k := -1; k <= 1; k++ {
				if j+k >= 0 && j+k < len(A[i]) {
					A[i][j] = min(A[i][j], temp+A[i-1][j+k])
				}
			}
		}
	}
	sort.Ints(A[len(A)-1])
	return A[len(A)-1][0]
}
