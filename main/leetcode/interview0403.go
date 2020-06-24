package main

func listOfDepth(tree *TreeNode) []*ListNode {
	var res []*ListNode
	if tree == nil {
		return res
	}
	queue := []*TreeNode{tree}
	for len(queue) != 0 {
		head := &ListNode{Val: queue[0].Val}
		res = append(res, head)
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			if i != 0 {
				head.Next = &ListNode{Val: node.Val}
				head = head.Next
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return res
}
