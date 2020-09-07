package main

//递归式写法
func preorderTraversal(root *TreeNode) []int {
	var res []int

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			res = append(res, root.Val)
			dfs(root.Left)
			dfs(root.Right)
		}
	}

	dfs(root)
	return res
}

//迭代式(?)写法
func preorderTraversal2(root *TreeNode) []int {
	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root)
	}
	var res []int

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
