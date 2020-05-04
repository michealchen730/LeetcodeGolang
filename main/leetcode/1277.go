package main

import "fmt"

func countSquares(matrix [][]int) int {
	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				res += getSquares(matrix, i, j)
			}
		}
	}
	return res
}

func getSquares(matrix [][]int, i, j int) int {
	res, size, row, col := 1, 1, len(matrix), len(matrix[0])
	lmt := min(row-i, col-j)
	for size < lmt {
		flag1, flag2 := true, true
		for m := i; m < i+size; m++ {
			if matrix[m][j+size] != 1 {
				flag1 = false
				break
			}
		}
		for k := j; k <= j+size; k++ {
			if matrix[i+size][k] != 1 {
				flag2 = false
				break
			}
		}
		if flag1 && flag2 {
			res++
			size++
		} else {
			break
		}
	}
	return res
}

func main() {
	fmt.Println(countSquares([][]int{[]int{0, 1, 1, 1}, []int{1, 1, 1, 1}, []int{0, 1, 1, 1}}))
}
