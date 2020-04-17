package main

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	mp := make(map[string]int)
	res := 0
	for _, v := range emails {
		var temp strings.Builder
		split := strings.Split(v, "@")
		for _, b := range split[0] {
			if b == '+' {
				break
			}
			if b != '.' {
				temp.WriteByte(byte(b))
			}
		}
		temp.WriteByte('@')
		temp.WriteString(split[1])
		str := temp.String()
		if _, ok := mp[str]; !ok {
			res++
			mp[str]++
		}
	}
	return res
}
