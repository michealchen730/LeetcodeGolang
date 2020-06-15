package main

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	var cpy strings.Builder
	cpy.WriteString(s1)
	cpy.WriteString(s1)
	return !(KMP(cpy.String(), s2) == -1)
}
