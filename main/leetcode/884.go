package main

import (
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	str1 := strings.Fields(A)
	str2 := strings.Fields(B)
	mp := make(map[string]int)
	for _, v := range str1 {
		mp[v]++
	}
	for _, v := range str2 {
		mp[v]++
	}
	var res []string
	for k, v := range mp {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
