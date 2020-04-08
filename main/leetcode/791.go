package main

import (
	"sort"
	"strings"
)

type t791 struct {
	bytes []byte
	Str   string
}

func (s t791) Len() int { return len(s.bytes) }
func (s t791) Less(i, j int) bool {
	return strings.IndexByte(s.Str, s.bytes[i]) < strings.IndexByte(s.Str, s.bytes[j])
}
func (s t791) Swap(i, j int) { s.bytes[i], s.bytes[j] = s.bytes[j], s.bytes[i] }

func customSortString(S string, T string) string {
	bs := []byte(T)
	sort.Sort(t791{
		bytes: bs,
		Str:   S,
	})
	return string(bs)
}
