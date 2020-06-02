package main

func maxVowels(s string, k int) int {
	i, j := 0, 0
	res := 0
	for j < k {
		if s[j] == 'a' || s[j] == 'o' || s[j] == 'e' || s[j] == 'u' || s[j] == 'i' {
			res++
		}
		j++
	}
	temp := res
	for j < len(s) {
		if s[j] == 'a' || s[j] == 'o' || s[j] == 'e' || s[j] == 'u' || s[j] == 'i' {
			temp++
		}
		if s[i] == 'a' || s[i] == 'o' || s[i] == 'e' || s[i] == 'u' || s[i] == 'i' {
			temp--
		}
		i++
		j++
		res = max(res, temp)
	}
	return res
}
