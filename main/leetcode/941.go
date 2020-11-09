package main

func validMountainArray(A []int) bool {
	if len(A) >= 3 {
		i, j := 0, len(A)-1
		for i+1 < len(A) && A[i] < A[i+1] {
			i++
		}
		for j-1 >= 0 && A[j-1] > A[j] {
			j--
		}
		if i == j && i != len(A)-1 && i != 0 {
			return true
		}
	}
	return false
}
