package main

import "strings"

func lengthOfLastWord(s string) int {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return 0
	} else {
		return len(fields[len(fields)-1])
	}
}
