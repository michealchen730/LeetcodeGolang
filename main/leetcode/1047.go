package main

func removeDuplicates1047(S string) string {
	var stack []byte
	for k, _ := range S {
		if k == 0 || len(stack) == 0 || S[k] != stack[len(stack)-1] {
			stack = append(stack, S[k])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
