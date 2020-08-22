package main

func levelOrderO33(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ordered := true
	res := [][]int{{root.Val}}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		l := len(queue)
		var arr []int
		for i := l - 1; i >= 0; i-- {
			if ordered {
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
					arr = append(arr, queue[i].Right.Val)
				}
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
					arr = append(arr, queue[i].Left.Val)
				}
			} else {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
					arr = append(arr, queue[i].Left.Val)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
					arr = append(arr, queue[i].Right.Val)
				}
			}
		}
		ordered = !ordered
		queue = queue[l:]
		if len(arr) == 0 {
			break
		}
		res = append(res, arr)
	}
	return res
}
