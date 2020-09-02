package main

import (
	"math"
	"strconv"
)

//栈加字符串反转
func calculate(s string) int {
	var stack []string
	temp := 0 //temp表示当前的数字
	pow := 0
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			temp += int(math.Pow10(pow)) * int(s[i]-'0')
			pow++
		} else {
			if i != len(s)-1 && (s[i+1] >= '0' && s[i+1] <= '9') {
				stack = append(stack, strconv.Itoa(temp))
			}
			temp, pow = 0, 0

			if s[i] == '(' {
				sum += evalExpr(&stack)
			} else {
				//也许预处理应该消除字符串中的空格
				if s[i] != ' ' {
					stack = append(stack, string(s[i]))
				}
			}

		}
	}
	//如果第一个字符是数字字符，那么记得加上
	if s[0] >= '0' && s[0] <= '9' {
		stack = append(stack, strconv.Itoa(temp))
	}
	sum = evalExpr(&stack)
	return sum
}

func evalExpr(expr *[]string) int {
	if len(*expr) == 0 {
		return 0
	}
	res, _ := strconv.Atoi((*expr)[len(*expr)-1])
	for len(*expr) > 0 && (*expr)[len(*expr)-1] != ")" {
		if (*expr)[len(*expr)-1] == "+" {
			*expr = (*expr)[:len(*expr)-1]
			t, _ := strconv.Atoi((*expr)[len(*expr)-1])
			res += t
		} else if (*expr)[len(*expr)-1] == "-" {
			*expr = (*expr)[:len(*expr)-1]
			t, _ := strconv.Atoi((*expr)[len(*expr)-1])
			res -= t
		} else {
			*expr = (*expr)[:len(*expr)-1]
		}
	}
	if len(*expr) > 0 {
		*expr = (*expr)[:len(*expr)-1]
	}
	*expr = append(*expr, strconv.Itoa(res))

	return res
}
