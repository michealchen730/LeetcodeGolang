package main

import "strings"

func toGoatLatin(sentence string) string {
	var t strings.Builder
	t.WriteByte('a')
	mp := map[byte]bool{
		'I': true,
		'i': true,
		'A': true,
		'a': true,
		'E': true,
		'e': true,
		'O': true,
		'o': true,
		'U': true,
		'u': true,
	}
	fields := strings.Fields(sentence)
	var res strings.Builder
	for idx, v := range fields {
		if mp[v[0]] {
			res.WriteString(v)
		} else {
			res.WriteString(v[1:])
			res.WriteByte(v[0])
		}
		res.WriteString("ma")
		res.WriteString(t.String())
		t.WriteByte('a')
		if idx != len(fields)-1 {
			res.WriteByte(' ')
		}
	}
	return res.String()
}
