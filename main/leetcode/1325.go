package main

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	mp := make(map[*TreeNode]*TreeNode)
	var dfs1325 func(node *TreeNode)
	dfs1325 = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			for node.Val == target && node.Left == nil && node.Right == nil {
				if mp[node] == nil {
					root = nil
					break
				}
				if mp[node].Left == node {
					mp[node].Left = nil
				}
				if mp[node].Right == node {
					mp[node].Right = nil
				}
				node = mp[node]
			}
			return
		}
		if node.Left != nil {
			mp[node.Left] = node
			dfs1325(node.Left)
		}
		if node.Right != nil {
			mp[node.Right] = node
			dfs1325(node.Right)
		}
	}
	dfs1325(root)
	return root
}
