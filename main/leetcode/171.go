package main

func titleToNumber(s string) int {
	t := 1
	res := 0
	for k := len(s) - 1; k >= 0; k-- {
		res += (int(s[k]-'A') + 1) * t
		t *= 26
	}
	return res
}
