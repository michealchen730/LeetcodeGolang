package main

func sumNumbers(root *TreeNode) int {
	var dfs func(root *TreeNode, path *[]int)
	res := 0
	dfs = func(root *TreeNode, path *[]int) {
		*path = append(*path, root.Val)
		if root.Left == nil && root.Right == nil {
			t := 0
			for i := 0; i < len(*path); i++ {
				t *= 10
				t += (*path)[i]
			}
			res += t
			*path = (*path)[:len(*path)-1]
			return
		}
		if root.Left != nil {
			dfs(root.Left, path)
		}
		if root.Right != nil {
			dfs(root.Right, path)
		}
		*path = (*path)[:len(*path)-1]
	}
	var path []int
	if root != nil {
		dfs(root, &path)
	}
	return res
}
