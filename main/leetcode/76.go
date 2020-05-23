package main

func minWindow(s string, t string) string {
	L, R := 0, len(s)+1
	mp1, mp2 := make(map[byte]int), make(map[byte]int)
	for k, _ := range t {
		mp1[t[k]]++
	}
	//方法内写函数，很有意思
	check := func() bool {
		for k, v := range mp1 {
			if mp2[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < len(s); r++ {
		if mp1[s[r]] > 0 {
			mp2[s[r]]++
			for check() && l <= r {
				if r-l+1 == len(t) {
					return s[l : r+1]
				}
				if r-l+1 < R-L {
					R, L = r+1, l
				}
				if mp1[s[l]] > 0 {
					mp2[s[l]]--
				}
				l++
			}
		}
	}
	if R == len(s)+1 {
		return ""
	} else {
		return s[L:R]
	}
}
