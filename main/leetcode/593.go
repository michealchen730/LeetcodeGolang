package main

import (
	"sort"
)

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	mp1 := make(map[int]int)

	mp1[(p1[0]-p2[0])*(p1[0]-p2[0])+(p1[1]-p2[1])*(p1[1]-p2[1])]++
	mp1[(p1[0]-p3[0])*(p1[0]-p3[0])+(p1[1]-p3[1])*(p1[1]-p3[1])]++
	mp1[(p1[0]-p4[0])*(p1[0]-p4[0])+(p1[1]-p4[1])*(p1[1]-p4[1])]++
	mp1[(p2[0]-p3[0])*(p2[0]-p3[0])+(p2[1]-p3[1])*(p2[1]-p3[1])]++
	mp1[(p2[0]-p4[0])*(p2[0]-p4[0])+(p2[1]-p4[1])*(p2[1]-p4[1])]++
	mp1[(p3[0]-p4[0])*(p3[0]-p4[0])+(p3[1]-p4[1])*(p3[1]-p4[1])]++

	if len(mp1) != 2 {
		return false
	}
	var arr []int
	for k, _ := range mp1 {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	if arr[0] == 0 {
		return false
	}

	return true
}
