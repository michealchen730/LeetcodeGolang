package main

func isBalanced(root *TreeNode) bool {
	if root != nil {
		if abs(getTreeDepth(root.Left)-getTreeDepth(root.Right)) <= 1 {
			return isBalanced(root.Left) && isBalanced(root.Right)
		} else {
			return false
		}
	} else {
		return true
	}
}

func getTreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return max(getTreeDepth(root.Left), getTreeDepth(root.Right)) + 1
	}
}
