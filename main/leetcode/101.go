package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric2(root.Left, root.Right)
}

func isSymmetric2(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	} else {
		if left.Val != right.Val {
			return false
		} else {
			return isSymmetric2(left.Left, right.Right) && isSymmetric2(left.Right, right.Left)
		}
	}
}
