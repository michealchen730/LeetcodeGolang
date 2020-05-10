package main

func maxLevelSum(root *TreeNode) int {
	if root != nil {
		mx, res, level, queue := root.Val, 1, 0, []*TreeNode{root}
		for len(queue) != 0 {
			level++
			sum := 0
			length := len(queue)
			for i := 0; i < length; i++ {
				sum += queue[i].Val
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			if sum > mx {
				mx = sum
				res = level
			}
			queue = queue[length:]
		}
		return res
	}
	return 0
}
