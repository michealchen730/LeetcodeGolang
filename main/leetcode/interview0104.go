package main

func canPermutePalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	mp := make(map[byte]int)
	for k, _ := range s {
		mp[s[k]]++
	}
	//flag是一个标记，记录mp中出现奇数次的元素数量
	flag := 0
	for _, v := range mp {
		if v%2 == 1 {
			flag++
		}
		if flag > 1 {
			return false
		}
	}
	return true
}
