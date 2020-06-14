package main

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	fr, sc := 1, 2
	n -= 2
	for n > 0 {
		fr, sc = sc, fr+sc
		n--
	}
	return sc
}
