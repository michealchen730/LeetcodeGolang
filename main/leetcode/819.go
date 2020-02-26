package main

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	reg := regexp.MustCompile("\\w+")
	mp1, mp2 := make(map[string]int), make(map[string]int)
	for _, v := range banned {
		mp1[v] = 1
	}
	if reg != nil {
		strs := reg.FindAllString(paragraph, -1)
		for _, v := range strs {
			if mp1[strings.ToLower(v)] != 1 {
				mp2[strings.ToLower(v)]++
			}
		}
	}
	max := 0
	res := ""
	for k, v := range mp2 {
		if v > max {
			max = v
			res = k
		}
	}
	return res
}
