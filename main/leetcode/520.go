package main

//规则解法
//func detectCapitalUse(word string) bool {
//	for k,_:=range word{
//		if word[k]<='Z'{
//			if k>0&&word[k-1]>='a'{
//				return false
//			}
//		}else{
//			if k>1&&word[k-1]<='Z'{
//				return false
//			}
//		}
//	}
//	return true
//}

//状态机解法
//此题完全可以当作状态机的例题，难度一般，入门再合适不过了
func detectCapitalUse(word string) bool {
	//FSM[0]代表起始状态，遇见A-Z进入FSM[1]，遇见a-z进入FSM[2]，
	//FSM[1]代表第一个字母大写，遇见A-Z进入FSM[3]，遇见a-z进入FSM[2]，
	//FSM[2]代表当前字母小写，遇见A-Z进入错误状态，遇见a-z进入FSM[2]，
	//FSM[3]代表当前字符串全部为大写且字符串长度超过两个，遇见A-Z进入FSM[3]，遇见a-z进入错误状态，
	FSM := [][]int{{1, 2}, {3, 2}, {-1, 2}, {3, -1}}
	t := 0
	for k, _ := range word {
		if word[k] <= 'Z' {
			t = FSM[t][0]
		} else {
			t = FSM[t][1]
		}
		if t == -1 {
			return false
		}
	}
	return true
}
