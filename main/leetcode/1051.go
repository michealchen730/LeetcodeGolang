package main

import (
	"sort"
)

func heightChecker(heights []int) int {
	if len(heights)<=1 {
		return len(heights)
	}
	temp:=make([]int,len(heights))
	copy(temp,heights)
	sort.Ints(temp)
	res:=0
	for i:=0;i<len(heights);i++{
		if temp[i]!=heights[i] {
			res++
		}
	}
	return res
}