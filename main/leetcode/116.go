package main

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

func connect(root *Node116) *Node116 {
	if root != nil {
		queue := []*Node116{root}
		for len(queue) != 0 {
			layer := len(queue)
			for i := 0; i < layer; i++ {
				if i < layer-1 {
					queue[i].Next = queue[i+1]
				}
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			queue = queue[layer:]
		}
	}
	return root
}
