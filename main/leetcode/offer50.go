package main

func firstUniqChar(s string) byte {
	mp := make(map[byte]int)
	for k, _ := range s {
		mp[s[k]]++
	}
	for k, _ := range s {
		if mp[s[k]] == 1 {
			return s[k]
		}
	}
	return ' '
}
