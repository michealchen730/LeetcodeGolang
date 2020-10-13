package main

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0

	var dfs func(root *TreeNode, flag int)

	dfs = func(root *TreeNode, flag int) {
		if flag == 1 && root.Left == nil && root.Right == nil {
			res += root.Val
		}

		if root.Left != nil {
			dfs(root.Left, 1)
		}

		if root.Right != nil {
			dfs(root.Right, 2)
		}

	}

	if root != nil {
		dfs(root, 2)
	}

	return res

}
