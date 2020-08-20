package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	inc, temp := true, 0
	arr := make([]strings.Builder, numRows)
	for k, _ := range s {
		arr[temp].WriteByte(s[k])
		if inc {
			temp++
			if temp == numRows-1 {
				inc = false
			}
		} else {
			temp--
			if temp == 0 {
				inc = true
			}
		}
	}
	for j := 1; j < numRows; j++ {
		arr[0].WriteString(arr[j].String())
	}
	return arr[0].String()
}
