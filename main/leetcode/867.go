package main

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	if m == n {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
		return matrix
	} else {
		res := make([][]int, n)
		for k, _ := range res {
			res[k] = make([]int, m)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				res[i][j] = matrix[j][i]
			}
		}
		return res
	}
}
