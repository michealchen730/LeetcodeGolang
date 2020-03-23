package main

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1[0] == '0' || num2[0] == '0' {
		return "0"
	}
	bytes1 := []byte(num1)
	bytes2 := []byte(num2)
	res := "0"
	s := ""
	for i := len(num2) - 1; i >= 0; i-- {
		var temp strings.Builder
		temp.WriteString(s)
		s += "0"
		n := 0
		t := int(bytes2[i] - '0')
		if t != 0 {
			for j := len(num1) - 1; j >= 0; j-- {
				mu := t*int(bytes1[j]-'0') + n
				temp.WriteString(strconv.Itoa(mu % 10))
				n = mu / 10
				if j == 0 && n != 0 {
					temp.WriteString(strconv.Itoa(n))
				}
			}
		}
		res = plusStrings(reverseStrings(temp.String()), res)
	}
	return res
}

func plusStrings(a, b string) string {
	i, j := len(a)-1, len(b)-1
	var res strings.Builder
	n := 0
	for i >= 0 && j >= 0 {
		t1, _ := strconv.Atoi(a[i : i+1])
		t2, _ := strconv.Atoi(b[j : j+1])
		t3 := t1 + t2 + n
		res.WriteString(strconv.Itoa(t3 % 10))
		n = t3 / 10
		i--
		j--
	}
	if i != j {
		for i >= 0 {
			t1, _ := strconv.Atoi(a[i : i+1])
			t3 := t1 + n
			res.WriteString(strconv.Itoa(t3 % 10))
			n = t3 / 10
			i--
		}
		for j >= 0 {
			t1, _ := strconv.Atoi(b[j : j+1])
			t3 := t1 + n
			res.WriteString(strconv.Itoa(t3 % 10))
			n = t3 / 10
			j--
		}
	}
	if n != 0 {
		res.WriteString(strconv.Itoa(n))
	}
	return reverseStrings(res.String())
}

func reverseStrings(s string) string {
	bytes := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
