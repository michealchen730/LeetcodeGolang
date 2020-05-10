package main

//先看dp解法
func mctFromLeafValues(arr []int) int {
	//数据的预处理
	matrix := make([][]int, len(arr))
	for k, _ := range matrix {
		matrix[k] = make([]int, len(arr))
	}
	for k := len(matrix) - 1; k >= 0; k-- {
		t := arr[k]
		for i := k; i < len(matrix); i++ {
			t = max(t, arr[i])
			matrix[k][i] = t
		}
	}

	//dp
	dp := make([][]int, len(arr))
	for k, _ := range dp {
		dp[k] = make([]int, len(arr))
	}
	for i := len(dp) - 1; i >= 0; i-- {
		for j := i + 1; j < len(dp); j++ {
			dp[i][j] = dp[i][i] + dp[i+1][j] + matrix[i][i]*matrix[i+1][j]
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+matrix[i][k]*matrix[k+1][j])
			}
		}
	}
	return dp[0][len(dp)-1]
}

//单调栈解法
func mctFromLeafValues2(arr []int) int {
	var stack []int
	res := 0
	for _, v := range arr {
		for len(stack) != 0 && stack[len(stack)-1] < v { //这里要使用for，举个例子[6 3 2]，v为4
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//这里注意栈为空的情况
			if len(stack) == 0 {
				res += t * v
				continue
			}
			res += t * min(stack[len(stack)-1], v)
		}
		//单调栈，一定会压栈的
		stack = append(stack, v)
	}
	if len(stack) >= 2 {
		//如果有这种情况，说明是一个单调栈，举个例子[6 5 4 3 2 1]
		//这时候需要处理这个栈
		for k := len(stack) - 2; k >= 0; k-- {
			res += stack[k] * stack[k+1]
		}
	}
	return res
}
