package main

func pathSum113(root *TreeNode, sum int) [][]int {
	var res [][]int
	var path []int
	var dfs func(root *TreeNode, temp int, path *[]int)

	dfs = func(root *TreeNode, temp int, path *[]int) {

		temp += root.Val
		*path = append(*path, root.Val)

		if root.Left == nil && root.Right == nil {
			if temp == sum {
				cpy := make([]int, len(*path))
				copy(cpy, *path)
				res = append(res, cpy)
				*path = (*path)[:len(*path)-1]
			}
			return
		}

		if root.Left != nil {
			dfs(root.Left, temp, path)
		}

		if root.Right != nil {
			dfs(root.Right, temp, path)
		}

		*path = (*path)[:len(*path)-1]

	}

	dfs(root, 0, &path)

	return res
}
