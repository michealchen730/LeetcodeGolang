package main

func findTilt(root *TreeNode) int {
	res := 0
	var dfs func(r *TreeNode) int
	dfs = func(node *TreeNode) int {
		l, r := 0, 0
		if node.Left != nil {
			l = dfs(node.Left)
		}
		if node.Right != nil {
			r = dfs(node.Right)
		}
		res += abs(l - r)
		return node.Val + l + r
	}
	if root != nil {
		dfs(root)
	}
	return res
}
