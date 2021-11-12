package main

import (
	"strconv"
	"strings"
)

func getHint(secret string, guess string) string {
	count := 0
	mp1 := make(map[byte]int)
	mp2 := make(map[byte]int)
	for k, _ := range secret {
		if secret[k] == guess[k] {
			count++
			continue
		}
		mp1[secret[k]]++
		mp2[guess[k]]++
	}
	count2 := 0
	for k, v := range mp1 {
		count2 += min(v, mp2[k])
	}
	var res strings.Builder
	res.WriteString(strconv.Itoa(count))
	res.WriteString("A")
	res.WriteString(strconv.Itoa(count2))
	res.WriteString("B")
	return res.String()
}
