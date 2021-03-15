package main

//func isSymmetric(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	return isSymmetric2(root.Left, root.Right)
//}

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

//迭代写法
func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root, root}
	for len(queue) != 0 {
		n1, n2 := queue[0], queue[1]
		if n1 == nil && n2 == nil {
			queue = queue[2:]
			continue
		}
		if (n1 != nil && n2 == nil) || (n1 == nil && n2 != nil) || n1.Val != n2.Val {
			return false
		}
		queue = append(queue, n1.Left)
		queue = append(queue, n2.Right)
		queue = append(queue, n1.Right)
		queue = append(queue, n2.Left)
		queue = queue[2:]
	}
	return true
}
