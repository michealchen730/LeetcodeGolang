package main

//比较菜的模拟写法
func minInsertions1541(s string) int {
	cnt := 0
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if len(stack) > 0 && stack[len(stack)-1] == ')' {
				cnt++
				stack = stack[:len(stack)-2]
			}
			stack = append(stack, s[i])
		} else {
			//如果是右括号
			//1、如果此时栈为空，那么必然要加一个'('，并且')'入栈
			if len(stack) == 0 {
				stack = append(stack, '(')
				cnt++
				stack = append(stack, ')')
			} else {
				if stack[len(stack)-1] == ')' {
					stack = stack[:len(stack)-2]
				} else {
					stack = append(stack, ')')
				}
			}
		}
	}
	if len(stack) != 0 {
		//类似这种用例："(((()"
		if stack[len(stack)-1] == ')' {
			cnt++
			stack = stack[:len(stack)-2]
		}
	}
	//类似这种用例，"(((("
	return cnt + len(stack)*2
}

func minInsertions1542(s string) int {
	numLeft := 0
	continuesRight := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			//这种例子："()("
			if continuesRight != 0 {
				if numLeft == 0 {
					res += 2
					numLeft++
				} else {
					res += 1
				}
				continuesRight = 0
			} else {
				numLeft++
			}
		} else {
			if continuesRight == 1 {
				if numLeft == 0 {
					res += 1
				} else {
					numLeft--
				}
				continuesRight--
			} else {
				continuesRight++
			}
		}
	}
	if continuesRight != 0 {
		res += 1
		if numLeft == 0 {
			res++
		} else {
			numLeft--
		}
	}

	return res + numLeft*2

}

func minInsertions1543(s string) int {
	numLeft := 0
	continuesRight := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			//这种例子："()("
			if continuesRight != 0 {
				if numLeft == 0 {
					res += 2
					numLeft++
				} else {
					res += 1
				}
				continuesRight = 0
			} else {
				numLeft++
			}
		} else {
			if continuesRight == 1 {
				if numLeft == 0 {
					res += 1
				} else {
					numLeft--
				}
				continuesRight--
			} else {
				continuesRight++
			}
		}
	}
	if continuesRight != 0 {
		res += 1
		if numLeft == 0 {
			res++
		} else {
			numLeft--
		}
	}

	return res + numLeft*2

}
