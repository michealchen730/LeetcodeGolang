package main

func minDepth(root *TreeNode) int {
	if root == nil {
		//左右子结点全部为空（叶子结点）
		return 0
	} else if !(root.Left != nil && root.Right != nil) {
		//左右子结点只有一个为空
		if root.Left != nil {
			return minDepth(root.Left) + 1
		} else {
			return minDepth(root.Right) + 1
		}
	} else {
		//左右子结点都不为空
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}
