package main

func removeKdigits(num string, k int) string {
	if num == "0" || len(num) <= k {
		return "0"
	}
	var res []byte

	start, n := 0, k
	var stack []int
	for k < len(num) {
		for i := start; i <= k; i++ {
			if len(stack) == 0 {
				stack = append(stack, i)
				continue
			}
			for len(stack) > 0 && num[stack[len(stack)-1]] > num[i] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
		if !(len(res) == 0 && num[stack[0]] == '0') {
			res = append(res, num[stack[0]])
		}

		start = stack[len(stack)-1] + 1
		stack = stack[1:]
		k++
		n--
	}
	if len(res) == 0 {
		return "0"
	}
	return string(res)
}
