package main

func bstToGst(root *TreeNode) *TreeNode {
	dfs1038(root, []int{0})
	return root
}

//二叉树后序遍历
func dfs1038(root *TreeNode, arr []int) {
	if root.Right != nil {
		dfs1038(root.Right, arr)
	}
	arr[0] += root.Val
	root.Val = arr[0]
	if root.Left != nil {
		dfs1038(root.Left, arr)
	}
}
