package main

import "fmt"

func diameterOfBinaryTree(root *TreeNode) int {
	arr := []int{0}
	fmt.Println(maxDepth(root, arr))
	return arr[0]
}

func maxDepth(root *TreeNode, arr []int) int {
	if root == nil {
		return -1
	}
	depth1 := maxDepth(root.Left, arr) + 1
	depth2 := maxDepth(root.Right, arr) + 1
	if depth1+depth2 > arr[0] {
		arr[0] = depth2 + depth1
	}
	return max(depth1, depth2)
}
