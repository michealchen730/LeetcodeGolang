package main

func maxDepthAfterSplit(seq string) []int {
	var res, stack []int
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			if len(stack) != 0 {
				if stack[len(stack)-1] == 0 {
					stack = append(stack, 1)
					res = append(res, 1)
					continue
				}
			}
			stack = append(stack, 0)
			res = append(res, 0)
		} else {
			res = append(res, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

//优化写法，发现效率太低了
func maxDepthAfterSplit2(seq string) []int {
	var res []int
	if len(seq) == 0 {
		return res
	}
	res = append(res, 0)
	for i := 1; i < len(seq); i++ {
		if seq[i] != seq[i-1] {
			res = append(res, res[len(res)-1])
		} else {
			if res[len(res)-1] == 0 {
				res = append(res, 1)
			} else {
				res = append(res, 0)
			}
		}
	}
	return res
}
