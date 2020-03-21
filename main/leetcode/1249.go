package main

func minRemoveToMakeValid(s string) string {
	var stack []int
	bytes := []byte(s)
	i := 0
	for i < len(bytes) {
		if bytes[i] == '(' {
			stack = append(stack, i)
		}
		if bytes[i] == ')' {
			if len(stack) == 0 {
				bytes = append(bytes[0:i], bytes[i+1:]...)
				i--
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		i++
	}
	if len(stack) != 0 {
		for j := len(stack) - 1; j >= 0; j-- {
			bytes = append(bytes[0:stack[j]], bytes[stack[j]+1:]...)
		}
	}
	return string(bytes)
}
