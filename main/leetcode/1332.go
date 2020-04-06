package main

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return 2
		}
		i++
		j--
	}
	return 1
}
