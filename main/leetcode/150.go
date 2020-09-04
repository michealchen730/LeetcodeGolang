package main

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, v := range tokens {
		switch v {
		case "+":
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			ans := num1 + num2
			stack = stack[:len(stack)-2]
			stack = append(stack, ans)
		case "-":
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			ans := num2 - num1
			stack = stack[:len(stack)-2]
			stack = append(stack, ans)
		case "*":
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			ans := num2 * num1
			stack = stack[:len(stack)-2]
			stack = append(stack, ans)
		case "/":
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			ans := num2 / num1
			stack = stack[:len(stack)-2]
			stack = append(stack, ans)
		default:
			temp, _ := strconv.Atoi(v)
			stack = append(stack, temp)
		}
	}
	return stack[0]
}
