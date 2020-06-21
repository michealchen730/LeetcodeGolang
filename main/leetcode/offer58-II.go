package main

import "strings"

func reverseLeftWords(s string, n int) string {
	var res strings.Builder
	res.WriteString(s[n:])
	res.WriteString(s[0:n])
	return res.String()
}
