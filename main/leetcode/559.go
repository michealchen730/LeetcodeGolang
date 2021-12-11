package main

func maxDepth559(root *Node) int {
	depth := 0
	if root == nil {
		return 0
	}
	for _, v := range root.Children {
		depth = max(depth, maxDepth559(v))
	}
	return depth + 1
}
