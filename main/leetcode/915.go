package main

func partitionDisjoint(A []int) int {
	n := len(A)
	B := make([]int, n+1)
	B[n] = A[n-1]
	for i := n - 1; i >= 0; i-- {
		B[i] = min(B[i+1], A[i])
	}
	mx := A[0]
	for i := 0; i < len(A); i++ {
		mx = max(mx, A[i])
		if mx <= B[i+1] {
			return i + 1
		}
	}
	return -1
}
