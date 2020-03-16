package main

import (
	"strconv"
	"strings"
)

func compressString(S string) string {
	if len(S) <= 2 {
		return S
	}
	var res strings.Builder
	for i := 0; i < len(S); i++ {
		res.WriteByte(S[i])
		temp := 1
		for j := i + 1; j < len(S); j++ {
			if S[j] != S[i] {
				break
			} else {
				i++
				temp++
			}
		}
		res.WriteString(strconv.Itoa(temp))
	}
	if res.Len() >= len(S) {
		return S
	}
	return res.String()
}
