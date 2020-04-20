package main

//想要不排序，还是需要用双指针
func sortedSquares(A []int) []int {
	var res []int
	if A[len(A)-1] <= 0 {
		for i := len(A) - 1; i >= 0; i-- {
			res = append(res, A[i]*A[i])
		}
		return res
	}
	if A[0] >= 0 {
		for _, v := range A {
			res = append(res, v*v)
		}
		return res
	}
	if A[0] < 0 && A[len(A)-1] > 0 {
		i, j := 0, len(A)-1
		for i <= j {
			m1, m2 := A[i]*A[i], A[j]*A[j]
			res = append([]int{max(m1, m2)}, res...)
			if m1 < m2 {
				j--
			} else {
				i++
			}
		}
	}
	return res
}
