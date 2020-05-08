package main

import "sort"

func luckyNumbers(matrix [][]int) []int {
	var res []int
	for i := 0; i < len(matrix[0]); i++ {
		max := matrix[0][i]
		temp := 0
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > max {
				temp = j
				max = matrix[j][i]
			}
		}
		arr := make([]int, len(matrix[0]))
		copy(arr, matrix[temp])
		sort.Ints(arr)
		if arr[0] == max {
			res = append(res, max)
		}
	}
	return res
}
