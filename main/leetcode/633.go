package main

func judgeSquareSum(c int) bool {
	mp := make(map[int]int)
	for i := 0; i <= c; i++ {
		temp := i * i
		if temp < c {
			mp[temp]++
			if mp[c-temp] != 0 {
				return true
			}
		} else if temp == c {
			return true
		} else {
			break
		}
	}
	return false
}
