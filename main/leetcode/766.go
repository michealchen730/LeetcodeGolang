package main

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) <= 1 {
		return true
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
