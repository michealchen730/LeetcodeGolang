package main

import "strings"

func shortestPalindrome(s string) string {
	for j := len(s); j >= 0; j-- {
		if isPalindrome(s[:j]) {
			var sb strings.Builder
			for t := len(s) - 1; t >= j; t-- {
				sb.WriteByte(s[t])
			}
			return sb.String() + s
		}
	}
	return s
}
