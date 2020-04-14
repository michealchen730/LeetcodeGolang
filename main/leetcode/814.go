package main

func pruneTree(root *TreeNode) *TreeNode {
	if root.Left != nil {
		root.Left = pruneTree(root.Left)
	}
	if root.Right != nil {
		root.Right = pruneTree(root.Right)
	}
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	} else {
		return root
	}
}
