package main

import (
	"strconv"
	"strings"
)

func thousandSeparator(n int) string {
	if n < 1000 {
		return strconv.Itoa(n)
	}

	var res []byte
	if n == 0 {
		return "0"
	}

	d := 0
	for n > 0 {
		d++
		b := n % 10
		n = n / 10
		res = append(res, byte(b+'0'))
		if d%3 == 0 && n != 0 {
			res = append(res, '.')
		}

	}

	var sb strings.Builder
	for i := len(res) - 1; i >= 0; i-- {
		sb.WriteByte(res[i])
	}
	return sb.String()
}
