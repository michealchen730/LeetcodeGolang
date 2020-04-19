package main

func distributeCoins(root *TreeNode) int {
	ans := []int{0}
	dfs979(root, &ans)
	return ans[0]
}

func dfs979(root *TreeNode, ans *[]int) int {
	if root == nil {
		return 0
	}
	L := dfs979(root.Left, ans)
	R := dfs979(root.Right, ans)
	(*ans)[0] += abs(L) + abs(R)
	return root.Val + L + R - 1
}
