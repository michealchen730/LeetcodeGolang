package main

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	mp := make(map[int]int)
	used := make([]int, len(pushed))
	for k, v := range pushed {
		mp[v] = k
	}
	for i := 0; i < len(popped); i++ {
		//栈为空或者当前栈顶元素在pushed数组中的位置要小于当前popped元素时，发起进栈操作
		if len(stack) == 0 || mp[stack[len(stack)-1]] < mp[popped[i]] {
			j := 0
			for ; pushed[j] != popped[i]; j++ {
				if used[j] == 0 {
					stack = append(stack, pushed[j])
					used[j] = 1
				}
			}
			//假装popped元素进栈又出栈了
			used[j] = 2
			continue
		}
		//如果当前元素成功匹配栈顶，那么出栈
		if stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		//如果当前元素不匹配，返回false
		if mp[popped[i]] < mp[stack[len(stack)-1]] {
			return false
		}
	}
	return true
}

//思路进一步优化 (1)取消used数组，使用哨兵代替
func validateStackSequences2(pushed []int, popped []int) bool {
	var stack []int
	mp, t := make(map[int]int), 0
	for k, v := range pushed {
		mp[v] = k
	}
	for _, v := range popped {
		//栈为空或者当前栈顶元素在pushed数组中的位置要小于当前popped元素时，发起进栈操作
		if len(stack) == 0 || mp[stack[len(stack)-1]] < mp[v] {
			for ; pushed[t] != v; t++ {
				stack = append(stack, pushed[t])
			}
			//假装popped元素进栈又出栈了
			t++
			continue
		}
		//如果当前元素成功匹配栈顶，那么出栈
		if stack[len(stack)-1] == v {
			stack = stack[:len(stack)-1]
			continue
		}
		//如果当前元素不匹配，返回false
		if mp[v] < mp[stack[len(stack)-1]] {
			return false
		}
	}
	return true
}

//官方题解思路（妙啊）
func validateStackSequences3(pushed []int, popped []int) bool {
	var stack []int
	j := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) != 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return j == len(popped)
}
