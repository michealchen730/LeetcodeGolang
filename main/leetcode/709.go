package main

import "strings"

func toLowerCase(s string) string {
	var res strings.Builder
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			res.WriteRune(rune(v + 32))
		} else {
			res.WriteRune(v)
		}
	}
	return res.String()
}
