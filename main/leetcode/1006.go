package main

func clumsy(N int) int {
	stack := []int{N}
	opr := 0
	t := N - 1
	for t >= 1 {
		if opr == 0 {
			stack[len(stack)-1] *= t
		}
		if opr == 1 {
			stack[len(stack)-1] /= t
		}
		if opr == 2 {
			stack = append(stack, t)
		}
		if opr == 3 {
			stack = append(stack, -t)
		}
		opr++
		t--
		opr %= 4
	}
	res := 0
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}
