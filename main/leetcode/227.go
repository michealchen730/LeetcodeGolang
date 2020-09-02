package main

import (
	"strconv"
	"strings"
)

func calculate227(s string) int {
	var stack []string
	var sb strings.Builder
	for k, _ := range s {
		if s[k] != ' ' {
			sb.WriteByte(s[k])
		}
	}
	s = sb.String()

	temp := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			temp *= 10
			temp += int(s[i] - '0')
		} else {
			stack = append(stack, strconv.Itoa(temp))
			temp = 0
			stack = append(stack, string(s[i]))
		}
	}
	stack = append(stack, strconv.Itoa(temp))

	for i := 0; i < len(stack); i++ {
		if stack[i] == "*" {
			if i-2 >= 0 {
				stack[i] = stack[i-2]
			} else {
				stack[i] = "+"
			}

			mul1, _ := strconv.Atoi(stack[i-1])
			mul2, _ := strconv.Atoi(stack[i+1])
			stack[i+1] = strconv.Itoa(mul1 * mul2)
			stack[i-1] = "0"
		}

		if stack[i] == "/" {
			if i-2 >= 0 {
				stack[i] = stack[i-2]
			} else {
				stack[i] = "+"
			}

			mul1, _ := strconv.Atoi(stack[i-1])
			mul2, _ := strconv.Atoi(stack[i+1])
			stack[i+1] = strconv.Itoa(mul1 / mul2)
			stack[i-1] = "0"
		}
	}

	res, _ := strconv.Atoi(stack[0])

	for i := 1; i < len(stack); i++ {
		if stack[i] != "+" && stack[i] != "-" {
			if stack[i-1] == "+" {
				temp, _ := strconv.Atoi(stack[i])
				res += temp
			}
			if stack[i-1] == "-" {
				temp, _ := strconv.Atoi(stack[i])
				res -= temp
			}
		}
	}

	return res
}
