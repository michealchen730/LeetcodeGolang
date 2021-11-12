package main

//中序遍历
func kthSmallest230(root *TreeNode, k int) int {
	count := 0
	res := -1
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		count++
		if count == k {
			res = r.Val
			return
		}
		dfs(r.Right)
	}
	dfs(root)
	return res
}
