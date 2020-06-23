package main

import (
	"strings"
)

func addBinary(a string, b string) string {
	i, j, flag := len(a)-1, len(b)-1, 0
	var bytes []byte
	for i >= 0 && j >= 0 {
		t := int(a[i]-'0'+b[j]-'0') + flag
		if t >= 2 {
			flag = 1
			t = t % 2
		} else {
			flag = 0
		}
		bytes = append(bytes, byte(t+'0'))
		i--
		j--
	}
	for i >= 0 {
		t := int(a[i]-'0') + flag
		if t >= 2 {
			flag = 1
			t = t % 2

		} else {
			flag = 0
		}
		bytes = append(bytes, byte(t+'0'))
		i--
	}
	for j >= 0 {
		t := int(b[j]-'0') + flag
		if t >= 2 {
			flag = 1
			t = t % 2
		} else {
			flag = 0
		}
		bytes = append(bytes, byte(t+'0'))
		j--
	}
	if flag > 0 {
		bytes = append(bytes, '1')
	}

	var res strings.Builder
	for k := len(bytes) - 1; k >= 0; k-- {
		res.WriteByte(bytes[k])
	}
	return res.String()
}
