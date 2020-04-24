package main

func repeatedNTimes(A []int) int {
	for i := 1; i <= 3; i++ {
		for j := 0; j < len(A)-i; j++ {
			if A[j] == A[j+i] {
				return A[j]
			}
		}
	}
	return 0
}
