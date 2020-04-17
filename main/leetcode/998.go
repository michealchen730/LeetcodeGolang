package main

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root.Val < val {
		node := &TreeNode{
			Val:  val,
			Left: root,
		}
		return node
	}
	head := root
	for root.Val > val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
			return head
		} else {
			if root.Right.Val < val {
				node := &TreeNode{
					Left: root.Right,
					Val:  val,
				}
				root.Right = node

			} else {
				root = root.Right
			}
		}
	}
	return head
}
