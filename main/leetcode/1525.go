package main

func numSplits(s string) int {
	mp := make(map[byte]int)
	for k, _ := range s {
		mp[s[k]]++
	}
	mp2 := make(map[byte]int)
	res := 0
	for k, _ := range s {
		mp2[s[k]]++
		mp[s[k]]--
		if mp[s[k]] == 0 {
			delete(mp, s[k])
		}
		if len(mp) == len(mp2) {
			res++
		}
		if len(mp2) > len(mp) {
			break
		}
	}
	return res
}
