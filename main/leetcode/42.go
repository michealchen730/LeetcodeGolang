package main

func trap(height []int) int {
	var stack []int
	res := 0
	for i := 0; i < len(height); i++ {
		if height[i] != 0 {
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				temp := 0
				for len(stack) > 0 {
					if height[i] >= height[stack[len(stack)-1]] {
						res += (i - stack[len(stack)-1] - 1) * (height[stack[len(stack)-1]] - temp)
						temp = height[stack[len(stack)-1]]
						stack = stack[:len(stack)-1]
					} else {
						res += (height[i] - temp) * (i - stack[len(stack)-1] - 1)
						break
					}
				}
				stack = append(stack, i)
			}
		}
	}
	return res
}
