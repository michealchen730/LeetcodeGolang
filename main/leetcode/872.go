package main

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var arr1, arr2 []int
	dfs872(root1, &arr1)
	dfs872(root2, &arr2)
	if len(arr1) != len(arr2) {
		return false
	} else {
		for k, v := range arr1 {
			if arr2[k] != v {
				return false
			}
		}
		return true
	}
}

func dfs872(root *TreeNode, nodes *[]int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			*nodes = append(*nodes, root.Val)
			return
		}
		if root.Left != nil {
			dfs872(root.Left, nodes)
		}
		if root.Right != nil {
			dfs872(root.Right, nodes)
		}
	}
}
