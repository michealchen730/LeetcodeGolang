package main

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	mp1, mp2 := make(map[byte]int), make(map[byte]int)
	for k, _ := range s1 {
		mp1[s1[k]]++
		mp2[s2[k]]++
	}
	for k, v := range mp1 {
		if v != mp2[k] {
			return false
		}
	}
	return true
}
