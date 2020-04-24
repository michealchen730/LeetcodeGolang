package main

func isUnivalTree(root *TreeNode) bool {
	return dfs965(root, root.Val)
}
func dfs965(root *TreeNode, key int) bool {
	if root == nil {
		return true
	} else {
		return root.Val == key && dfs965(root.Left, key) && dfs965(root.Right, key)
	}
}
