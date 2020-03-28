package main

import (
	"sort"
	"strings"
)

type Words []string

func (s Words) Len() int           { return len(s) }
func (s Words) Less(i, j int) bool { return len(s[i]) > len(s[j]) }
func (s Words) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func minimumLengthEncoding(words []string) int {
	res := 0
	sort.Sort(Words(words))
	tree := Trie{}
	for _, v := range words {
		tmp := reBuildstr(v)
		if !tree.StartsWith(tmp) {
			tree.Insert(tmp)
			res += len(v) + 1
		}
	}
	return res
}

func reBuildstr(word string) string {
	var res strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		res.WriteByte(word[i])
	}
	return res.String()
}
