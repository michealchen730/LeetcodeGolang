package main

import "fmt"

func diameterOfBinaryTree(root *TreeNode) int {
	arr := []int{0}
	fmt.Println(maxDepth543(root, arr))
	return arr[0]
}

func maxDepth543(root *TreeNode, arr []int) int {
	if root == nil {
		return -1
	}
	depth1 := maxDepth543(root.Left, arr) + 1
	depth2 := maxDepth543(root.Right, arr) + 1
	if depth1+depth2 > arr[0] {
		arr[0] = depth2 + depth1
	}
	return max(depth1, depth2)
}
