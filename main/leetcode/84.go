package main

//第一种思路，简单的二重循环，以i（数组下标）作为这个矩形的最后一部分，往前搜索，更新最大值
func largestRectangleArea(heights []int) int {
	res := 0
	for k, v := range heights {
		mn, t := v, k
		for t >= 0 {
			mn = min(mn, heights[t])
			res = max(mn*(k-t+1), res)
			t--
		}
	}
	return res
}

//第二种思路，单调栈
func largestRectangleArea2(heights []int) int {
	//这里的思路比较高级
	res := 0
	var stack []int
	stack = append(stack, -1)
	for k, v := range heights {
		for len(stack) > 1 && v < heights[stack[len(stack)-1]] {
			t := stack[len(stack)-1]
			h := heights[t]
			stack = stack[:len(stack)-1]
			res = max(res, h*(k-stack[len(stack)-1]-1))
		}
		stack = append(stack, k)
	}
	l := len(heights) - 1
	for len(stack) > 1 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = max(res, heights[t]*(l-stack[len(stack)-1]))
	}
	return res
}
