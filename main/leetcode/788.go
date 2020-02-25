package main

import (
	"strconv"
	"strings"
)

func rotatedDigits(N int) int {
	res := 0
	bed := []string{"3", "4", "7"}
	good := []string{"2", "5", "6", "9"}
	for i := 1; i <= N; i++ {
		s := strconv.Itoa(i)
		flag := -1
		for _, v := range bed {
			if strings.Contains(s, v) {
				flag = 0
				break
			}
		}
		if flag == 0 {
			continue
		}
		for _, v := range good {
			if strings.Contains(s, v) {
				res++
				break
			}
		}
	}
	return res
}
