package main

//这题应该是既有DFS和BFS
//这里只关注DFS写法
func deepestLeavesSum(root *TreeNode) int {
	res := 0
	depth := 0
	var dfs1302 func(node *TreeNode, depth int)
	dfs1302 = func(node *TreeNode, dep int) {
		if node == nil {
			return
		}
		dep++
		if node.Left == nil && node.Right == nil {
			if dep > depth {
				depth = dep
				res = node.Val
			} else if dep == depth {
				res += node.Val
			}
		}
		dfs1302(node.Left, dep)
		dfs1302(node.Right, dep)
	}
	dfs1302(root, 0)
	return res
}
