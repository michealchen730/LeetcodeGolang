package main

func rightSideView(root *TreeNode) []int {
	var res []int
	var nodes []*TreeNode
	if root != nil {
		nodes = append(nodes, root)
		for len(nodes) != 0 {
			length := len(nodes)
			res = append(res, nodes[length-1].Val)
			i := 0
			for ; i < length; i++ {
				node := nodes[i]
				if node.Left != nil {
					nodes = append(nodes, node.Left)
				}
				if node.Right != nil {
					nodes = append(nodes, node.Right)
				}
			}
			nodes = nodes[i:]
		}
	}
	return res
}
