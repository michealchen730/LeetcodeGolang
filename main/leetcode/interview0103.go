package main

import "strings"

func replaceSpaces(S string, length int) string {
	var res strings.Builder
	for k := 0; k < length; k++ {
		if S[k] != ' ' {
			res.WriteByte(S[k])
		} else {
			res.WriteString("%20")
		}
	}
	return res.String()
}
