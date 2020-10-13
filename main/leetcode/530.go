package main

func getMinimumDifference(root *TreeNode) int {
	arr := inorderTraversal(root)
	res := arr[1] - arr[0]
	for j := 2; j < len(arr); j++ {
		res = min(res, arr[j]-arr[j-1])
	}
	return res
}
