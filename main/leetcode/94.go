package main

//递归写法
func inorderTraversal(root *TreeNode) []int {
	var res []int
	dfs94(root, &res)
	return res
}

func dfs94(root *TreeNode, res *[]int) {
	if root != nil {
		dfs94(root.Left, res)
		*res = append(*res, root.Val)
		dfs94(root.Right, res)
	}
}
