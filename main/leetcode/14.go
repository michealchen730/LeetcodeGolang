package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	Compare := func(str1, str2 string) string {
		i := 0
		for ; i < min(len(str1), len(str2)); i++ {
			if str1[i] != str2[i] {
				break
			}
		}
		return str1[0:i]
	}
	temp, next := strs[0], 1
	for next < len(strs) {
		temp = Compare(temp, strs[next])
		if temp == "" {
			return temp
		}
		next++
	}
	return temp
}
