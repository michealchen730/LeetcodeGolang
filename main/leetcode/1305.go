package main

import "fmt"

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1, arr2 := morrisVal(root1), morrisVal(root2)
	arr3 := make([]int, len(arr1)+len(arr2))
	fmt.Println(arr1)
	fmt.Println(arr2)
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr3[k] = arr1[i]
			i++
		} else {
			arr3[k] = arr2[j]
			j++
		}
		k++
	}
	if i == len(arr1) {
		for j < len(arr2) {
			arr3[k] = arr2[j]
			j++
			k++
		}
	}
	if j == len(arr1) {
		for i < len(arr1) {
			arr3[k] = arr1[i]
			i++
			k++
		}
	}
	return arr3
}

func morrisVal(root *TreeNode) []int {
	var res []int
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			pre := morrisGetPredecessor(root)
			if pre.Right == nil {
				pre.Right = root
				root = root.Left
			} else if pre.Right == root {
				pre.Right = nil
				res = append(res, root.Val)
				root = root.Right
			}
		}
	}
	return res
}

func morrisGetPredecessor(root *TreeNode) *TreeNode {
	node := root
	if root.Left != nil {
		root = root.Left
		for root.Right != nil && root.Right != node {
			root = root.Right
		}
	}
	return root
}
