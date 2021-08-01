package main

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var res []int
	mp := make(map[*TreeNode]*TreeNode)
	used := make(map[*TreeNode]bool)
	var traverse func(r *TreeNode)
	traverse = func(r *TreeNode) {
		if r.Left != nil {
			mp[r.Left] = r
			traverse(r.Left)
		}
		if r.Right != nil {
			mp[r.Right] = r
			traverse(r.Right)
		}
	}
	traverse(root)
	var dfs func(r *TreeNode, k int)
	dfs = func(r *TreeNode, k int) {
		used[r] = true
		if k == 0 {
			res = append(res, r.Val)
			return
		}
		if r.Left != nil && !used[r.Left] {
			dfs(r.Left, k-1)
		}
		if r.Right != nil && !used[r.Right] {
			dfs(r.Right, k-1)
		}
		if mp[r] != nil && !used[mp[r]] {
			dfs(mp[r], k-1)
		}
	}
	dfs(target, k)
	return res
}
