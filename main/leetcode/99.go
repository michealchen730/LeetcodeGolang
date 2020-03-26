package main

func recoverTree(root *TreeNode) {
	//还是应该理解为找逆序对，使用中序遍历
	var pre, x, y *TreeNode
	for root != nil {
		if root.Left != nil {
			node := root.Left
			for node.Right != nil && node.Right != root {
				node = node.Right
			}
			if node.Right == nil {
				node.Right = root
				root = root.Left
			} else {
				if pre != nil && root.Val < pre.Val {
					y = root
					if x == nil {
						x = pre
					}
				}
				pre = root
				node.Right = nil
				root = root.Right
			}
		} else {
			if pre != nil && root.Val < pre.Val {
				y = root
				if x == nil {
					x = pre
				}
			}
			pre = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}
