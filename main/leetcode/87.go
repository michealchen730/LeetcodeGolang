package main

//比较熟悉递归，先用递归的方式写
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	mp := make(map[byte]int)
	for k, _ := range s1 {
		mp[s1[k]]++
		mp[s2[k]]--
	}
	for _, v := range mp {
		if v != 0 {
			return false
		}
	}

	for i := 1; i < len(s1); i++ {
		if (isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:])) || (isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i])) {
			return true
		}
	}
	return false

}
