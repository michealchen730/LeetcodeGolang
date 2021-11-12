package main

import (
	"strings"
)

//利用题解思路
func removeInvalidParentheses(s string) []string {
	//算出应该去除的左括号和右括号数目
	lCount, rCount := validParentheses(s)
	delete := make([]bool, len(s))
	hashtable := make(map[string]bool)
	var dfs func(lLeft, rLeft int, idx int, str string)
	dfs = func(lLeft, rLeft int, idx int, str string) {
		if lLeft == 0 && rLeft == 0 {
			var ans strings.Builder
			for k, v := range str {
				if !delete[k] {
					ans.WriteRune(v)
				}
			}
			l, r := validParentheses(ans.String())
			if l == 0 && r == 0 {
				hashtable[ans.String()] = true
			}
			return
		}
		if idx > len(str)-(rLeft+lLeft) {
			return
		}
		//对于每个左括号和右括号都可以选择要或者不要
		//对于其他元素直接不要
		dfs(lLeft, rLeft, idx+1, str)
		if idx == 0 || str[idx] != str[idx-1] || (idx > 0 && str[idx] == str[idx-1] && delete[idx-1]) {
			delete[idx] = true
			if str[idx] == '(' && lLeft > 0 {
				dfs(lLeft-1, rLeft, idx+1, str)
			}
			if str[idx] == ')' && rLeft > 0 {
				dfs(lLeft, rLeft-1, idx+1, str)
			}
			delete[idx] = false
		}
	}
	dfs(lCount, rCount, 0, s)
	var res []string
	for k, _ := range hashtable {
		res = append(res, k)
	}
	return res
}

func validParentheses(s string) (int, int) {
	lCount, rCount := 0, 0
	for _, v := range s {
		if v == '(' {
			lCount++
		} else if v == ')' {
			if lCount == 0 {
				rCount++
			} else {
				lCount--
			}
		}
	}
	return lCount, rCount
}
