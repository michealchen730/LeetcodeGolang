package main

import (
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	mp := make(map[byte]int)
	licensePlate = strings.ToLower(licensePlate)
	for i := 0; i < len(licensePlate); i++ {
		if 'a' <= licensePlate[i] && 'z' >= licensePlate[i] {
			mp[licensePlate[i]]++
		}
	}
	least := 0
	for _, v := range mp {
		least += v
	}
	index := -1
	for i := 0; i < len(words); i++ {
		if len(words[i]) < least {
			continue
		}
		if index != -1 && len(words[i]) >= len(words[index]) {
			continue
		}
		length := 0
		for k, v := range mp {
			if strings.Count(words[i], string(k)) < v {
				break
			} else {
				length++
			}
		}
		if length == len(mp) {
			index = i
		}
	}
	if index == -1 {
		return ""
	}
	return words[index]
}
