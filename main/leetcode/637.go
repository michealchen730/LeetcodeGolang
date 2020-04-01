package main

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root != nil {
		stack := []*TreeNode{root}
		for len(stack) != 0 {
			level := len(stack)
			num := 0
			for i := 0; i < level; i++ {
				num += stack[i].Val
				if stack[i].Left != nil {
					stack = append(stack, stack[i].Left)
				}
				if stack[i].Right != nil {
					stack = append(stack, stack[i].Right)
				}
			}
			res = append(res, float64(num)/float64(level))
			stack = stack[level:]
		}
	}
	return res
}
