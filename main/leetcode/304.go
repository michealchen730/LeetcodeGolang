package main

type NumMatrix struct {
	mat [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	mat := make([][]int, len(matrix))
	for k, _ := range mat {
		mat[k] = make([]int, len(matrix[0]))
		tmp := 0
		for idx, _ := range mat[k] {
			tmp += matrix[k][idx]
			mat[k][idx] = tmp
			if k > 0 {
				mat[k][idx] += mat[k-1][idx]
			}

		}
	}
	return NumMatrix{
		mat: mat,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := this.mat[row2][col2]
	if row1 != 0 {
		res -= this.mat[row1-1][col2]
	}
	if col1 != 0 {
		res -= this.mat[row2][col1-1]
	}
	if row1 != 0 && col1 != 0 {
		res += this.mat[row1-1][col1-1]
	}
	return res
}
