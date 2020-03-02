package main

func checkRecord(s string) bool {
	if len(s) <= 1 {
		return true
	}
	flag1 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			flag1++
			if flag1 == 2 {
				return false
			}
		}
		if i >= 2 && s[i] == 'L' {
			if s[i-1] == 'L' && s[i-2] == 'L' {
				return false
			}
		}
	}
	return true
}
