package main

import (
	"sort"
)

func orderlyQueue(S string, K int) string {
	bytes := []byte(S)
	if K >= 2 {
		sort.Sort(Bytes(bytes))
		return string(bytes)
	}
	if K == 1 {
		for i := 1; i < len(S); i++ {
			bytes = append(bytes, bytes[0])
			bytes = bytes[1:]
			if compareStr(bytes, S) {
				S = string(bytes)
			}
		}
	}
	return S
}

func compareStr(b []byte, s string) bool {
	for i := 0; i < len(b); i++ {
		if b[i]-'a' < s[i]-'a' {
			return true
		} else if b[i]-'a' > s[i]-'a' {
			return false
		}
	}
	return false
}
