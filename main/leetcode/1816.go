package main

import (
	"strings"
)

func truncateSentence(s string, k int) string {
	var res strings.Builder
	count := 0
	for _, v := range s {
		if v == ' ' {
			count++
			if count == k {
				return res.String()
			}
		}
		res.WriteRune(v)
	}
	return s
}
