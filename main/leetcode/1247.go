package main

func minimumSwap(s1 string, s2 string) int {
	type1, type2 := 0, 0
	for k, _ := range s1 {
		if s1[k] != s2[k] {
			if s1[k] == 'x' {
				type1++
			} else {
				type2++
			}
		}
	}
	if (type1+type2)%2 != 0 {
		return -1
	} else {
		return type1/2 + type2/2 + (type1%2)*2
	}
}
