package main

import "strings"

func restoreString(s string, indices []int) string {
	mp := make(map[int]byte)
	for k, _ := range s {
		mp[indices[k]] = s[k]
	}
	var res strings.Builder
	for i := 0; i < len(indices); i++ {
		res.WriteByte(mp[i])
	}
	return res.String()
}
