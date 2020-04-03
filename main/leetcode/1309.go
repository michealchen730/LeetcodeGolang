package main

import (
	"strconv"
	"strings"
)

func freqAlphabets(s string) string {
	var res strings.Builder
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && s[i+2] == '#' {
			temp, _ := strconv.Atoi(s[i : i+2])
			res.WriteByte(byte('a' + temp - 1))
			i += 2
		} else {
			temp, _ := strconv.Atoi(s[i : i+1])
			res.WriteByte(byte('a' + temp - 1))
		}
	}
	return res.String()
}
