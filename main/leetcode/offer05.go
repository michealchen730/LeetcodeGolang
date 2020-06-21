package main

import "strings"

func replaceSpace(s string) string {
	var res strings.Builder
	for _, v := range s {
		if v == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteByte(byte(v))
		}
	}
	return res.String()
}
