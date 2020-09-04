package main

import "fmt"

func searchMatrix240(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	mn, mx := 0, len(matrix[0])-1
	//mn,mx表示允许搜索的最小范围和最大范围
	for i := 0; i < len(matrix); i++ {
		low, high := mn, mx
		for low <= high {
			mid := low + (high-low)/2
			if matrix[i][mid] > target {
				high = mid - 1
			} else if matrix[i][mid] < target {
				low = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix240([][]int{[]int{5}, []int{6}}, 6))
}
