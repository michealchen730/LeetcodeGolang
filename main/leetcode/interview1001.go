package main

func merge2(A []int, m int, B []int, n int) {
	if len(B) == 0 {
		return
	}
	p1, p2, p3 := m-1, n-1, m+n-1
	for p1 >= 0 && p2 >= 0 {
		if A[p1] <= B[p2] {
			A[p3] = B[p2]
			p2--
		} else {
			A[p3] = A[p1]
			p1--
		}
		p3--
	}
	if p1 < 0 {
		for i := 0; i <= p2; i++ {
			A[i] = B[i]
		}
	}
}
