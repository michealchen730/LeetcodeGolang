package main

import (
	"strings"
)

func longestNiceSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}
	for k := 0; k <= 25; k++ {
		t1, t2 := strings.IndexByte(s, byte('a'+k)), strings.IndexByte(s, byte('A'+k))
		mx, mn := max(t1, t2), min(t1, t2)
		if mx != -1 && mn == -1 {
			res1, res2 := longestNiceSubstring(s[:mx]), longestNiceSubstring(s[mx+1:])
			if len(res1) > len(res2) {
				return res1
			} else {
				return res2
			}
		}
	}
	return s
}
