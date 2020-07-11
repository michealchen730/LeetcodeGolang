package main

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	var a1, a2, b1, b2 int
	for i := 0; i < len(a); i++ {
		if a[i] == '+' {
			if a[0] != '-' {
				a1, _ = strconv.Atoi(a[:i])
			} else {
				a1, _ = strconv.Atoi(a[1:i])
				a1 = -a1
			}
			break
		}
	}
	for i := 0; i < len(b); i++ {
		if b[i] == '+' {
			if b[0] != '-' {
				b1, _ = strconv.Atoi(b[:i])
			} else {
				b1, _ = strconv.Atoi(b[1:i])
				b1 = -b1
			}
			break
		}
	}
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] == '+' {
			if a[i+1] != '-' {
				a2, _ = strconv.Atoi(a[i+1 : len(a)-1])
			} else {
				a2, _ = strconv.Atoi(a[i+2 : len(a)-1])
				a2 = -a2
			}
		}
	}
	for i := len(b) - 2; i >= 0; i-- {
		if b[i] == '+' {
			if b[i+1] != '-' {
				b2, _ = strconv.Atoi(b[i+1 : len(b)-1])
			} else {
				b2, _ = strconv.Atoi(b[i+2 : len(b)-1])
				b2 = -b2
			}
		}
	}
	var res strings.Builder
	res.WriteString(strconv.Itoa(a1*b1 - a2*b2))
	res.WriteByte('+')
	res.WriteString(strconv.Itoa(a1*b2 + b1*a2))
	res.WriteByte('i')
	return res.String()
}
