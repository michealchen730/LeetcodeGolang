package main

func isBalanced110(node *TreeNode) bool {
	return isBalancedHelper(node) != -1
}

func isBalancedHelper(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := isBalancedHelper(node.Left) //!左节点
	if l == -1 {                     //剪枝，不平衡时直接返回，
		return -1
	}

	r := isBalancedHelper(node.Right) //!右节点
	if r == -1 {                      //剪枝，不平衡时直接返回
		return -1
	}

	if abs(l-r) > 1 { //剪枝，不平衡时直接返回
		return -1
	}

	return max(l, r) + 1 //计算深度 !根节点
}
