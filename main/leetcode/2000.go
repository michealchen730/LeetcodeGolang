package main

import (
	"strings"
)

func reversePrefix(word string, ch byte) string {
	t := strings.IndexByte(word, ch)
	if t != -1 {
		return reverseStr2(word, t)
	}
	return word
}

func reverseStr2(s string, k int) string {
	if len(s) <= 1 {
		return s
	}
	bytes := []byte(s)
	swapBytes(bytes, 0, k)
	return string(bytes)
}
