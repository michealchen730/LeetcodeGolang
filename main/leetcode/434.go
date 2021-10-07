package main

import "strings"

func countSegments(s string) int {
	fields := strings.Fields(s)
	return len(fields)
}
