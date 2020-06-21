package main

import "math"

func maxPathSum(root *TreeNode) int {
	arr := []int{math.MinInt32}
	maxGain(root, arr)
	return arr[0]
}

func maxGain(node *TreeNode, arr []int) int {
	if node == nil {
		return 0
	}
	// 递归计算左右子节点的最大贡献值
	// 只有在最大贡献值大于 0 时，才会选取对应子节点
	leftGain := max(maxGain(node.Left, arr), 0)
	rightGain := max(maxGain(node.Right, arr), 0)

	// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
	priceNewPath := node.Val + leftGain + rightGain

	// 更新答案
	arr[0] = max(arr[0], priceNewPath)

	// 返回节点的最大贡献值
	return node.Val + max(leftGain, rightGain)
}
