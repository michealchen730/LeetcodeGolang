package main

import (
	"sort"
	"strings"
)

type S767 struct {
	b byte
	c int
}

type S767s []S767

func (s S767s) Len() int           { return len(s) }
func (s S767s) Less(i, j int) bool { return s[i].c > s[j].c }
func (s S767s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func reorganizeString(S string) string {
	arr := make([]int, 26)
	for _, v := range S {
		arr[v-'a']++
	}
	mx := 0
	Strs := make([]S767, 0, 26)
	for k, v := range arr {
		if v != 0 {
			Strs = append(Strs, S767{
				b: byte(k + 'a'),
				c: v,
			})
		}
		mx = max(mx, v)
	}
	if mx > (len(S)+1)/2 {
		return ""
	}

	var res strings.Builder
	l := 0
	for l < len(S) {
		sort.Sort(S767s(Strs))
		res.WriteByte(Strs[0].b)
		Strs[0].c--
		l++
		if l < len(S) {
			res.WriteByte(Strs[1].b)
			Strs[1].c--
			l++
		}
	}

	return res.String()
}
