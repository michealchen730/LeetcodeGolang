package main

import "strings"

func hasAllCodes(s string, k int) bool {
	temp := 1 << k
	for i := 0; i < temp; i++ {
		if !strings.Contains(s, toBits(i, k)) {
			return false
		}
	}
	return true
}

func toBits(k, l int) string {
	var b []byte
	if k == 0 {
		b = append(b, '0')
	}
	for k != 0 {
		b = append([]byte{byte(48 + k%2)}, b...)
		k = k / 2
	}
	for len(b) < l {
		b = append([]byte{'0'}, b...)
	}
	return string(b)
}
