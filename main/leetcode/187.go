package main

func findRepeatedDnaSequences(s string) []string {
	var res []string
	mp := make(map[string]int)
	for i := 9; i < len(s); i++ {
		mp[s[i-9:i+1]]++
		if mp[s[i-9:i+1]] == 2 {
			res = append(res, s[i-9:i+1])
		}
	}
	return res
}
