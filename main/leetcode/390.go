package main

//递归写法
//找规律
//创建映射关系
func lastRemaining390(n int) int {
	if n == 1 {
		return 1
	}
	return 2 * (n/2 + 1 - lastRemaining390(n/2))
}
