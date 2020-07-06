package main

//递归做法
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	if getDepth(root.Left) >= getDepth(root.Right) {
		return findBottomLeftValue(root.Left)
	} else {
		return findBottomLeftValue(root.Right)
	}
}

//DFS+BFS做法

func findBottomLeftValue2(root *TreeNode) int {
	depth := getDepth(root)
	if depth == 1 {
		return root.Val
	}
	queue := []*TreeNode{root}
	depth--
	for len(queue) != 0 {
		depth--
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				if depth == 0 {
					return node.Left.Val
				} else {
					queue = append(queue, node.Left)
				}
			}
			if node.Right != nil {
				if depth == 0 {
					return node.Right.Val
				} else {
					queue = append(queue, node.Right)
				}
			}
		}
		queue = queue[l:]
	}
	return -1
}
