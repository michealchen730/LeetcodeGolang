package main

import "strings"

func reverseParentheses(s string) string {
	var stack []int
	bytes := []byte(s)
	for k, v := range s {
		if v == '(' {
			stack = append(stack, k)
		}
		if v == ')' {
			left := stack[len(stack)-1]
			reverseBytes1190(bytes, left, k)
			stack = stack[:len(stack)-1]
		}
	}
	var res strings.Builder
	for _, v := range bytes {
		if v != '(' && v != ')' {
			res.WriteByte(v)
		}
	}
	return res.String()
}

func reverseBytes1190(arr []byte, i int, k int) {
	for i < k {
		arr[i], arr[k] = arr[k], arr[i]
		i++
		k--
	}
}
