package main

func findDiagonalOrder(matrix [][]int) []int {
	var res, temp []int
	row, col, flag, i, j := len(matrix), len(matrix[0]), 0, 0, 0
	if row == 1 {
		return matrix[0]
	}
	for j <= col-1 {
		m, n := i, j
		temp = append(temp, matrix[i][j])
		for m != 0 && n != col-1 {
			m--
			n++
			temp = append(temp, matrix[m][n])
		}
		if flag%2 != 0 {
			reverseArr(temp)
		}
		res = append(res, temp...)
		flag++
		temp = temp[:0]
		if i != row-1 {
			i++
		} else {
			j++
		}
	}
	return res
}
