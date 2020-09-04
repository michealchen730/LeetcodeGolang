package main

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {

	var path []int
	var res []string

	var dfs func(path []int, root *TreeNode)
	dfs = func(path []int, root *TreeNode) {
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			var s strings.Builder
			for i := 0; i < len(path); i++ {
				s.WriteString(strconv.Itoa(path[i]))
				if i != len(path)-1 {
					s.WriteString("->")
				}
			}
			res = append(res, s.String())
			path = path[:len(path)-1]
			return
		}
		if root.Left != nil {
			dfs(path, root.Left)
		}
		if root.Right != nil {
			dfs(path, root.Right)
		}

		path = path[:len(path)-1]
	}
	if root != nil {
		dfs(path, root)
	}
	return res
}
