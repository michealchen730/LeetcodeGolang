package main

import (
	"sort"
)

type Mat1289 [][]int

func (s Mat1289) Len() int           { return len(s) }
func (s Mat1289) Less(i, j int) bool { return s[i][0] < s[j][0] }
func (s Mat1289) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func minFallingPathSum(arr [][]int) int {
	length := len(arr)
	for i := 0; i < len(arr)-1; i++ {
		var mat [][]int
		for k, v := range arr[i] {
			mat = append(mat, []int{v, k})
		}
		sort.Sort(Mat1289(mat))
		for j := 0; j < len(arr[0]); j++ {
			for m := 0; m < len(mat); m++ {
				if mat[m][1] != j {
					arr[i+1][j] += mat[m][0]
					break
				}
			}
		}
	}
	sort.Ints(arr[len(arr)-1])
	return arr[length-1][0]
}
