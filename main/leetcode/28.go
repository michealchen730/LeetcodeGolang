package main

func strStr28(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	i, j := 0, 0
	next := getNext28(needle)
	for i < len(haystack) {
		if j == len(needle)-1 && needle[j] == haystack[i] {
			return i - j
		}
		if j == -1 || needle[j] == haystack[i] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	return -1
}

func getNext28(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	i, j := 0, -1
	for i < len(s)-1 {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}
