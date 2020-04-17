package main

import (
	"sort"
)

type IntS2 [][]int

func (s IntS2) Len() int { return len(s) }
func (s IntS2) Less(i, j int) bool {
	if s[i][0] < s[j][0] {
		return true
	} else if s[i][0] == s[j][0] {
		return s[i][1] < s[j][1]
	} else {
		return false
	}
}
func (s IntS2) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func intervalIntersection(A [][]int, B [][]int) [][]int {
	var res [][]int
	A = append(A, B...)
	sort.Sort(IntS2(A))
	k := 0
	for i := 1; i < len(A); i++ {
		if A[i][0] <= A[k][1] {
			res = append(res, []int{A[i][0], min(A[k][1], A[i][1])})
		}
		if A[i][1] >= A[k][1] {
			k = i
		}
	}
	return res
}

//双指针写一下
func intervalIntersection2(A [][]int, B [][]int) [][]int {
	var res [][]int
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		if A[i][1] > B[j][1] {
			if A[i][0] <= B[j][1] {
				res = append(res, []int{max(A[i][0], B[j][0]), B[j][1]})
			}
			j++
		} else if A[i][1] == B[j][1] {
			res = append(res, []int{max(A[i][0], B[j][0]), A[i][1]})
			i++
			j++
		} else {
			if B[j][0] <= A[i][1] {
				res = append(res, []int{max(A[i][0], B[j][0]), A[i][1]})
			}
			i++
		}
	}
	return res
}
