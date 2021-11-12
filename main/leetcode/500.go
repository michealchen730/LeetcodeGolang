package main

import (
	"strings"
)

func findWords500(words []string) []string {
	keyboards := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	maps := []map[rune]bool{{}, {}, {}}
	for k, v := range keyboards {
		for _, b := range v {
			maps[k][b] = true
		}
	}
	var res []string
	for _, v := range words {
		ts := strings.ToLower(v)
	SE:
		for _, m := range maps {
			if !m[rune(ts[0])] {
				continue
			} else {
				for _, r := range ts {
					if !m[r] {
						break SE
					}
				}
				res = append(res, v)
			}
		}
	}
	return res
}
