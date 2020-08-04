package main

//与70题相同

func numWays(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	fr, sc := 1, 2
	n -= 2
	for n > 0 {
		fr, sc = sc%1000000007, (fr+sc)%1000000007
		n--
	}
	return sc
}
