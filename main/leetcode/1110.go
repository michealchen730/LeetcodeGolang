package main

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	//所有delete存进map
	mp := make(map[int]int)
	for _, v := range to_delete {
		mp[v]++
	}
	var res []*TreeNode
	//BFS
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[0]
		if mp[node.Val] == 0 {
			res = append(res, node)
			//BFS遍历
			queue := []*TreeNode{node}
			for len(queue) != 0 {
				lev := len(queue)
				for i := 0; i < lev; i++ {
					temp := queue[i]
					if temp.Left != nil {
						if mp[temp.Left.Val] != 0 {
							stack = append(stack, temp.Left)
							temp.Left = nil
						} else {
							queue = append(queue, temp.Left)
						}
					}
					if temp.Right != nil {
						if mp[temp.Right.Val] != 0 {
							stack = append(stack, temp.Right)
							temp.Right = nil
						} else {
							queue = append(queue, temp.Right)
						}
					}
				}
				queue = queue[lev:]
			}
		} else {
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		stack = stack[1:]
	}
	return res
}
