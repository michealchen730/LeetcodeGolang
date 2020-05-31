package main

type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}

func connect117(root *Node117) *Node117 {
	if root != nil {
		queue := []*Node117{root}
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
