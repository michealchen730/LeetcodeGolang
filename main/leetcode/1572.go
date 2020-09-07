package main

func diagonalSum(mat [][]int) int {
	res := 0
	for i := 0; i < len(mat); i++ {
		res += mat[i][i]
		if len(mat)-1-i != i {
			res += mat[i][len(mat)-1-i]
		}
	}
	return res
}
