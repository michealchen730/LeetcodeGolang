package main

func findClosest(words []string, word1 string, word2 string) int {
	mp := make(map[string][]int)
	for k, v := range words {
		mp[v] = append(mp[v], k)
	}
	a1, a2 := mp[word1], mp[word2]
	p1, p2 := 0, 0
	res := len(words)
	for p1 < len(a1) && p2 < len(a2) {
		res = min(res, abs(a1[p1]-a2[p2]))
		if a1[p1] < a2[p2] {
			p1++
		} else {
			p2++
		}
	}
	return res
}
