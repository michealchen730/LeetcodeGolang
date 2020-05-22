package main

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	h1, h2 := 0, 0
	if root.Left != nil {
		h1 = getDepth(root.Left)
	}
	if root.Right != nil {
		h2 = getDepth(root.Right)
	}
	if h1 > h2 {
		return subtreeWithAllDeepest(root.Left)
	} else if h1 < h2 {
		return subtreeWithAllDeepest(root.Right)
	} else {
		return root
	}
}
