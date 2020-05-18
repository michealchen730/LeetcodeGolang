package main

func convertBST(root *TreeNode) *TreeNode {
	dfs1038(root, []int{0})
	return root
}
