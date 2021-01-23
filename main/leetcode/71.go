package main

import (
	"strings"
)

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	var stack []string
	for _, v := range paths {
		if v == "" || v == "." {
			continue
		}
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	var res strings.Builder
	for _, v := range stack {
		res.WriteByte('/')
		res.WriteString(v)
	}
	return res.String()
}
