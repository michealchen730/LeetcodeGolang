package main

func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	mp := map[byte]bool{
		'a': true,
		'A': true,
		'o': true,
		'O': true,
		'e': true,
		'E': true,
		'u': true,
		'U': true,
		'I': true,
		'i': true,
	}
	bytes := []byte(s)
	for left < right {
		if !mp[s[left]] {
			left++
			continue
		}
		if !mp[s[right]] {
			right--
			continue
		}
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
	return string(bytes)
}
