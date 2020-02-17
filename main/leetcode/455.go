package main

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	if len(g)==0||len(s)==0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	i,j:=0,0
	res:=0
	for i<len(g)&&j<len(s){
		if s[j]>=s[i]  {
			res++
			i++
		}
		j++
	}
	return res
}
