package main

func findSecondMinimumValue(root *TreeNode) int {
	res := root.Val
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r != nil {
			if r.Val == root.Val {
				dfs(r.Left)
				dfs(r.Right)
			} else {
				if res == root.Val {
					res = r.Val
				} else {
					res = min(res, r.Val)
				}
			}
		}
	}
	dfs(root)
	if res == root.Val {
		return -1
	}
	return res
}
