package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := root
	insert := &TreeNode{Val: val}
	for node.Left != nil || node.Right != nil {
		if node.Val > val {
			if node.Left == nil {
				node.Left = insert
				return root
			} else {
				node = node.Left
			}
		} else {
			if node.Right == nil {
				node.Right = insert
				return root
			} else {
				node = node.Right
			}
		}
	}
	if node.Val > val {
		node.Left = insert
	} else {
		node.Right = insert
	}
	return root
}
