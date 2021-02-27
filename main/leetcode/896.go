package main

func isMonotonic(A []int) bool {
	flag := 0
	for k, v := range A {
		if k > 0 {
			if v > A[k-1] {
				if flag == -1 {
					return false
				} else {
					flag = 1
				}
			}
			if v < A[k-1] {
				if flag == 1 {
					return false
				} else {
					flag = -1
				}
			}
		}
	}
	return true
}
