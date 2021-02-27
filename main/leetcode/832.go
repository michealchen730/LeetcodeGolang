package main

func flipAndInvertImage(A [][]int) [][]int {
	//先水平翻转
	for k, v := range A {
		n := len(v)
		i, j := 0, n-1
		for i < j {
			A[k][i], A[k][j] = A[k][j]^1, A[k][i]^1
			i++
			j--
		}
	}
	return A
}
