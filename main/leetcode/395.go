package main

import "strings"

func longestSubstring(s string, k int) int {
	ans := 0
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}

	var split byte
	for i, c := range cnt {
		if c > 0 && c < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}

	for _, subStr := range strings.Split(s, string(split)) {
		ans = max(ans, longestSubstring(subStr, k))
	}
	return ans
}
