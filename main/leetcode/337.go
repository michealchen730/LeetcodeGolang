package main

func rob3(root *TreeNode) int {
	val := dfs337(root)
	return max(val[0], val[1])
}

func dfs337(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l, r := dfs337(root.Left), dfs337(root.Right)
	selected := l[1] + r[1] + root.Val
	notselected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notselected}
}
