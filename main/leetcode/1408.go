package main

import (
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	var res []string
	sort.Sort(StringsByLength(words))
	for i := len(words) - 2; i >= 0; i-- {
		for j := len(words) - 1; j >= i+1; j-- {
			if strings.Contains(words[j], words[i]) {
				res = append(res, words[i])
				break
			}
		}
	}
	return res
}
