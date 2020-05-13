package main

func sumRootToLeaf(root *TreeNode) int {
	var arr []int
	res := []int{0}
	if root != nil {
		dfs1022(root, &arr, res)
	}
	return res[0]
}

func dfs1022(root *TreeNode, arr *[]int, sum []int) {
	*arr = append(*arr, root.Val)
	if root.Left == nil && root.Right == nil {
		sum[0] += getSum1022(*arr)
		return
	}
	if root.Left != nil {
		dfs1022(root.Left, arr, sum)
	}
	if root.Right != nil {
		dfs1022(root.Right, arr, sum)
	}
	*arr = (*arr)[:len(*arr)-1]
}

func getSum1022(arr []int) int {
	res := 0
	temp := 1
	for i := len(arr) - 1; i >= 0; i-- {
		res += temp * arr[i]
		temp <<= 1
	}
	return res
}
