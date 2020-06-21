package main

import "sort"

func diagonalSort(mat [][]int) [][]int {
	for i := 0; i < len(mat); i++ {
		t := i
		var arr []int
		for j := 0; t < len(mat) && j < len(mat[0]); j++ {
			arr = append(arr, mat[t][j])
			t++
		}
		sort.Ints(arr)
		t = i
		for j := 0; t < len(mat) && j < len(mat[0]); j++ {
			mat[t][j] = arr[j]
			t++
		}
	}
	for i := 1; i < len(mat[0]); i++ {
		t := i
		var arr []int
		for j := 0; t < len(mat[0]) && j < len(mat); j++ {
			arr = append(arr, mat[j][t])
			t++
		}
		sort.Ints(arr)
		t = i
		for j := 0; t < len(mat[0]) && j < len(mat); j++ {
			mat[j][t] = arr[j]
			t++
		}
	}
	return mat
}
