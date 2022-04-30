package main

import (
	"fmt"
	"strings"
)

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	//枚举第一个数(最右边)的位置
	for i := 1; i <= len(num)-2; i++ {
		add1 := num[:i]
		//03这样的数是不被允许的
		if len(add1) > 1 && add1[0] == '0' {
			continue
		}
		for j := i + 1; j <= len(num)-1; j++ {
			add2 := num[i:j]
			if len(add2) > 1 && add2[0] == '0' {
				continue
			}
			//求和
			add3 := stringAdd(add1, add2)
			t := j
			for len(add3)+t <= len(num) && add3 == num[t:len(add3)+t] {
				fmt.Println(add2)
				fmt.Println(add3)
				if len(add3)+t == len(num) {
					return true
				}
				t = len(add3) + t
				add4 := stringAdd(add2, add3)
				add2 = add3
				add3 = add4
			}
		}
	}
	return false
}

func stringAdd(num1 string, num2 string) string {
	t1, t2 := len(num1), len(num2)
	var plus strings.Builder
	if t1 != t2 {
		for i := 0; i < abs(t1-t2); i++ {
			plus.WriteByte('0')
		}
		if t1 < t2 {
			plus.WriteString(num1)
			num1 = plus.String()
		} else {
			plus.WriteString(num2)
			num2 = plus.String()
		}
	}
	var res []byte
	carry := 0
	t := len(num1) - 1
	for t >= 0 {
		if int(num1[t]-'0'+num2[t]-'0')+carry >= 10 {
			res = append(res, byte(int(num1[t]-'0'+num2[t])+carry-10))
			carry = 1
		} else {
			res = append(res, byte(int(num1[t]-'0'+num2[t])+carry))
			carry = 0
		}
		t--
	}
	if carry > 0 {
		res = append(res, '1')
	}
	var ans strings.Builder
	for k := len(res) - 1; k >= 0; k-- {
		ans.WriteByte(res[k])
	}
	return ans.String()
}
