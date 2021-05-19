package main

func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Val >= low && r.Val <= high {
			res += r.Val
		}
		if r.Val > low {
			dfs(r.Left)
		}
		if r.Val < high {
			dfs(r.Right)
		}
	}
	dfs(root)
	return res
}
