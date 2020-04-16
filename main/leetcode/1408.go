package main

import (
	"sort"
	"strings"
)

type StringS []string

func (s StringS) Len() int           { return len(s) }
func (s StringS) Less(i, j int) bool { return len(s[i]) < len(s[j]) }
func (s StringS) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func stringMatching(words []string) []string {
	var res []string
	sort.Sort(StringS(words))
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
