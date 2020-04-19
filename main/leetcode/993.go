package main

func isCousins(root *TreeNode, x int, y int) bool {
	//表示x,y节点的父节点
	var pars []int
	if root != nil {
		queue := []*TreeNode{root}
		//标记位，当出现x,y值时增加1
		for len(queue) != 0 {
			flag := 0
			level := len(queue)
			for i := 0; i < level; i++ {
				if queue[i].Left != nil {
					temp := queue[i].Left.Val
					if temp == x || temp == y {
						flag++
						pars = append(pars, queue[i].Val)
					}
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					temp := queue[i].Right.Val
					if temp == x || temp == y {
						flag++
						pars = append(pars, queue[i].Val)
					}
					queue = append(queue, queue[i].Right)
				}
			}
			if flag == 2 {
				if pars[0] == pars[1] {
					return false
				} else {
					return true
				}
			} else if flag == 1 {
				return false
			}
			queue = queue[level:]
		}
	}
	return false
}
