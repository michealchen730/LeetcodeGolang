package main

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	for true {
		lDepth := getDepth(root.Left)
		rDepth := getDepth(root.Right)
		if lDepth > rDepth {
			root = root.Left
		} else if lDepth < rDepth {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
