package main

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if sum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
