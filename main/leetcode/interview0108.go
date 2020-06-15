package main

func setZeroes0108(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for k, v := range row {
		if v {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[k][j] = 0
			}
		}
	}
	for k, v := range col {
		if v {
			for i := 0; i < len(matrix); i++ {
				matrix[i][k] = 0
			}
		}
	}
}
