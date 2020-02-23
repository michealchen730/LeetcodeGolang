package main

import (
	"unicode"
)

func letterCasePermutation(S string) []string {
	runes := []rune(S)
	var res []string
	goGetPermutation(0, runes, &res)
	return res
}

func goGetPermutation(i int, runes []rune, res *[]string) {
	if i == len(runes) {
		s := string(runes)
		*res = append(*res, s)
		return
	}
	if 'a' <= runes[i] && runes[i] <= 'z' {
		goGetPermutation(i+1, runes, res)
		runes[i] = unicode.ToUpper(runes[i])
		goGetPermutation(i+1, runes, res)
	} else if 'A' <= runes[i] && runes[i] <= 'Z' {
		goGetPermutation(i+1, runes, res)
		runes[i] = unicode.ToLower(runes[i])
		goGetPermutation(i+1, runes, res)
	} else {
		goGetPermutation(i+1, runes, res)
	}
}
