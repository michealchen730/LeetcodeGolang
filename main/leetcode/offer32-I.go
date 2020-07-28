package main

func levelOrderO32I(root *TreeNode) []int {
	var res []int
	if root != nil {
		queue := []*TreeNode{root}
		for len(queue) != 0 {
			lev := len(queue)
			var level []int
			for i := 0; i < lev; i++ {
				res = append(res, queue[i].Val)
				level = append(level, queue[i].Val)
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			queue = queue[lev:]
		}
	}
	return res
}
