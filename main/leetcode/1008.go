package main

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}
	for i := 1; i < len(preorder); i++ {
		buildBST(preorder[i], root)
	}
	return root
}

func buildBST(i int, root *TreeNode) {
	if i < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: i}
			return
		} else {
			buildBST(i, root.Left)
		}
	}
	if i > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: i}
			return
		} else {
			buildBST(i, root.Right)
		}
	}
}
