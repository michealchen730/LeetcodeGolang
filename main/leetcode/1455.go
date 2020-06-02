package main

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	fields := strings.Fields(sentence)
	for k, v := range fields {
		if len(v) >= len(searchWord) {
			if v[0:len(searchWord)] == searchWord {
				return k + 1
			}
		}
	}
	return -1
}
