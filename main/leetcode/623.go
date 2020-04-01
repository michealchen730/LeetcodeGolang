package main

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		return &TreeNode{Left: root, Val: v}
	}
	if root != nil {
		stack := []*TreeNode{root}
		depth := 1
		for depth < d-1 {
			length := len(stack)
			for i := 0; i < length; i++ {
				if stack[i].Left != nil {
					stack = append(stack, stack[i].Left)
				}
				if stack[i].Right != nil {
					stack = append(stack, stack[i].Right)
				}
			}
			stack = stack[length:]
			depth++
		}
		for _, k := range stack {
			left, right := &TreeNode{Val: v, Left: k.Left}, &TreeNode{Val: v, Right: k.Right}
			k.Left, k.Right = left, right
		}
	}
	return root
}
