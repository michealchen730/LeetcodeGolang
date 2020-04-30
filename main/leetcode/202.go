package main

func isHappy(n int) bool {
	mp := make(map[int]int)
	for n != 1 {
		if _, ok := mp[n]; ok {
			return false
		} else {
			mp[n] = 1
		}
		sum := 0
		for n != 0 {
			t := n % 10
			sum += t * t
			n = n / 10
		}
		n = sum
	}
	return true
}
