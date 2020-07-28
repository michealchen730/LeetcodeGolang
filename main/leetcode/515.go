package main

func largestValues(root *TreeNode) []int {
	var res []int
	if root != nil {
		queue := []*TreeNode{root}
		for len(queue) != 0 {
			l := len(queue)
			mx := queue[0].Val
			for i := 0; i < l; i++ {
				node := queue[i]
				mx = max(mx, node.Val)
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
			res = append(res, mx)
			queue = queue[l:]
		}
	}
	return res
}
