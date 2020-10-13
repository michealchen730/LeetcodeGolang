package main

import "math"

func minCameraCover(root *TreeNode) int {
	var dfs func(root *TreeNode) (a, b, c int)

	dfs = func(root *TreeNode) (a, b, c int) {
		if root == nil {
			return math.MaxInt32, 0, 0
		}
		la, lb, lc := dfs(root.Left)
		ra, rb, rc := dfs(root.Right)
		a = lc + rc + 1
		b = min(a, min(la+rb, lb+ra))
		c = min(a, lb+rb)
		return
	}

	_, res, _ := dfs(root)
	return res
}
