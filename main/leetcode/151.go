package main

import "strings"

func reverseWords(s string) string {
	fields := strings.Fields(s)
	i, j := 0, len(fields)-1
	for i < j {
		fields[i], fields[j] = fields[j], fields[i]
		i++
		j--
	}
	var res strings.Builder
	for k, v := range fields {
		res.WriteString(v)
		if k != len(fields)-1 {
			res.WriteByte(' ')
		}
	}
	return res.String()
}
