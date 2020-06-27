package main

func maxDepthOffer55(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepthOffer55(root.Left), maxDepthOffer55(root.Right)) + 1
}
