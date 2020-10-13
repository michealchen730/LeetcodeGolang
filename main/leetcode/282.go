package main

import (
	"strconv"
	"strings"
)

//递归加回溯
//需要尝试所有的可能性
//这题的难点：
//1、加减法和乘法优先级
//2、数字拼接是否会溢出（题目的样例并未限制了num的长度）
func addOperators(num string, target int) []string {
	var res []string
	tar := int64(target)
	var ops []string

	if len(num) == 0 {
		return res
	}

	dfs282(num, tar, 0, 0, 0, 0, &ops, &res)
	return res
}

func dfs282(num string, tar int64, index int, previousOperand, currentOperand, value int64, ops, res *[]string) {

	if index == len(num) {
		if value == tar && currentOperand == 0 {
			var sb strings.Builder
			for _, v := range *ops {
				sb.WriteString(v)
			}
			*res = append(*res, sb.String())
		}
		return
	}

	currentOperand = currentOperand*10 + int64(num[index]-'0')
	current_val_rep := strconv.FormatInt(currentOperand, 10)

	if currentOperand > 0 {
		dfs282(num, tar, index+1, previousOperand, currentOperand, value, ops, res)
	}

	*ops = append(*ops, "+")
	*ops = append(*ops, current_val_rep)
	dfs282(num, tar, index+1, currentOperand, 0, value+currentOperand, ops, res)
	*ops = (*ops)[:len(*ops)-2]

	if len(*ops) > 0 {
		*ops = append(*ops, "-")
		*ops = append(*ops, current_val_rep)
		dfs282(num, tar, index+1, -currentOperand, 0, value-currentOperand, ops, res)
		*ops = (*ops)[:len(*ops)-2]

		*ops = append(*ops, "+")
		*ops = append(*ops, current_val_rep)
		dfs282(num, tar, index+1, currentOperand*previousOperand, 0, value-previousOperand+currentOperand*previousOperand, ops, res)
		*ops = (*ops)[:len(*ops)-2]

	}
}
