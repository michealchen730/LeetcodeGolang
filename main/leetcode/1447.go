package main

import (
	"strconv"
	"strings"
)

func simplifiedFractions(n int) []string {
	var res []string
	for t := 2; t <= n; t++ {
		for i := 1; i < t; i++ {
			if gcd(i, t) == 1 {
				var str strings.Builder
				s := strconv.Itoa(i)
				str.WriteString(s)
				str.WriteString("/")
				str.WriteString(strconv.Itoa(t))
				res = append(res, str.String())
			}
		}
	}
	return res
}

func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}
