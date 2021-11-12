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
//func addOperators(num string, target int) []string {
//	var res []string
//	tar := int64(target)
//	var ops []string
//
//	if len(num) == 0 {
//		return res
//	}
//
//	dfs282(num, tar, 0, 0, 0, 0, &ops, &res)
//	return res
//}
//
//func dfs282(num string, tar int64, index int, previousOperand, currentOperand, value int64, ops, res *[]string) {
//
//	if index == len(num) {
//		if value == tar && currentOperand == 0 {
//			var sb strings.Builder
//			for _, v := range *ops {
//				sb.WriteString(v)
//			}
//			*res = append(*res, sb.String())
//		}
//		return
//	}
//
//	currentOperand = currentOperand*10 + int64(num[index]-'0')
//	current_val_rep := strconv.FormatInt(currentOperand, 10)
//
//	if currentOperand > 0 {
//		dfs282(num, tar, index+1, previousOperand, currentOperand, value, ops, res)
//	}
//
//	*ops = append(*ops, "+")
//	*ops = append(*ops, current_val_rep)
//	dfs282(num, tar, index+1, currentOperand, 0, value+currentOperand, ops, res)
//	*ops = (*ops)[:len(*ops)-2]
//
//	if len(*ops) > 0 {
//		*ops = append(*ops, "-")
//		*ops = append(*ops, current_val_rep)
//		dfs282(num, tar, index+1, -currentOperand, 0, value-currentOperand, ops, res)
//		*ops = (*ops)[:len(*ops)-2]
//
//		*ops = append(*ops, "+")
//		*ops = append(*ops, current_val_rep)
//		dfs282(num, tar, index+1, currentOperand*previousOperand, 0, value-previousOperand+currentOperand*previousOperand, ops, res)
//		*ops = (*ops)[:len(*ops)-2]
//	}
//}

func addOperators(num string, target int) []string {
	targetInt64 := int64(target)
	var res []string
	op := []string{"*", "-", "+"}
	// num的currentPos后插入符号
	var dfs func(currentPos int, previousPos int, tmpRes int64, lastRes int64, ops []string, nums []string)
	dfs = func(currentPos int, previousPos int, tmpRes int64, lastRes int64, ops []string, nums []string) {
		//防止出现05这样的用例
		if num[previousPos] == '0' && currentPos != previousPos+1 {
			return
		}
		l := int64(0)
		numToHandle := num[previousPos:currentPos]
		numtmp, _ := strconv.Atoi(numToHandle)
		numTmp := int64(numtmp)
		nums = append(nums, numToHandle)
		// 表示还没有有效的ops
		if previousPos == 0 || ops[len(ops)-1] == "+" {
			tmpRes += numTmp
			l = numTmp
		} else if ops[len(ops)-1] == "-" {
			tmpRes -= numTmp
			l = -numTmp
		} else {
			l = lastRes * numTmp
			tmpRes = tmpRes - lastRes + lastRes*numTmp
		}

		//到达最后一位了
		if currentPos == len(num) {
			if tmpRes == targetInt64 {
				var ans strings.Builder
				ans.WriteString(nums[0])
				for k, v := range ops {
					ans.WriteString(v)
					ans.WriteString(nums[k+1])
				}
				res = append(res, ans.String())
			}
			return
		}

		for j := currentPos + 1; j <= len(num); j++ {
			for _, v := range op {
				ops = append(ops, v)
				dfs(j, currentPos, tmpRes, l, ops, nums)
				ops = ops[:len(ops)-1]
			}
		}
	}
	for i := 1; i <= len(num); i++ {
		dfs(i, 0, 0, 0, []string{}, []string{})
	}
	return res
}
