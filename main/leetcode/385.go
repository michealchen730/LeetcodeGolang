package main

import (
	"strconv"
	"unicode"
)

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool { return false }

func (n NestedInteger) GetInteger() int { return 0 }

func (n *NestedInteger) SetInteger(value int) {}

func (n *NestedInteger) Add(elem NestedInteger) {}

func (n NestedInteger) GetList() []*NestedInteger { return nil }

//这题不是我自己写的
func deserialize(s string) *NestedInteger {
	//s仅为数字，数字肯定是对的
	if s[0] != '[' {
		i, _ := strconv.Atoi(s)
		res := &NestedInteger{}
		res.SetInteger(i)
		return res
	}
	//否则肯定为列表
	stack := []*NestedInteger{}
	temp := 0
	neg := false
	for idx, v := range s {
		//找到嵌套
		if v == '[' {
			stack = append(stack, &NestedInteger{})
		} else if unicode.IsDigit(v) {
			temp = temp*10 + int(v-'0')
		} else if v == '-' {
			neg = true
		} else if v == ',' {
			//结算
			if unicode.IsDigit(rune(s[idx-1])) {
				if neg {
					temp = -temp
				}
				n := NestedInteger{}
				n.SetInteger(temp)
				stack[len(stack)-1].Add(n)
			}
			//将前置状态归零
			temp, neg = 0, false
		} else if v == ']' {
			//必然有一个NestedInteger结束
			//如果前一位是数字
			//那么这个NestedInteger完结
			if unicode.IsDigit(rune(s[idx-1])) {
				if neg {
					temp = -temp
				}
				n := NestedInteger{}
				n.SetInteger(temp)
				stack[len(stack)-1].Add(n)
			}
			//不是数字，那么就是']'
			//说明上一个NestedInteger已经完成
			//完成当前NestedInteger
			if len(stack) > 1 {
				stack[len(stack)-2].Add(*stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}
	return stack[0]
}
