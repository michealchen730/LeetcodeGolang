package main

import "fmt"

//二维前缀和
func maxSideLength(mat [][]int, threshold int) int {
	for i := 1; i < len(mat); i++ {
		mat[i][0] += mat[i-1][0]
	}
	for j := 1; j < len(mat[0]); j++ {
		mat[0][j] += mat[0][j-1]
	}
	for i := 1; i < len(mat); i++ {
		for j := 1; j < len(mat[0]); j++ {
			mat[i][j] = mat[i-1][j] + mat[i][j-1] - mat[i-1][j-1] + mat[i][j]
		}
	}
	row, col := len(mat), len(mat[0])
	left, right := 1, min(row, col)
	ans := 0
	for left <= right {
		mid, flag := left+(right-left)/2, true
	FI:
		for i := mid - 1; i < row; i++ {
			for j := mid - 1; j < col; j++ {
				sum := mat[i][j]
				if i != mid-1 && j != mid-1 {
					sum = sum - mat[i][j-mid] - mat[i-mid][j] + mat[i-mid][j-mid]
				}
				if i == mid-1 && j != mid-1 {
					sum -= mat[i][j-mid]
				}
				if j == mid-1 && i != mid-1 {
					sum -= mat[i-mid][j]
				}
				if sum <= threshold {
					fmt.Println(mid)

					flag = false
					ans = max(ans, mid)
					left = mid + 1
					break FI
				}
			}
		}
		if flag {
			right = mid - 1
		}
	}
	return ans
}
func main() {
	maxSideLength([][]int{[]int{1, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1}}, 5)
}
