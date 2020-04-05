package main

func oddEvenList(head *ListNode) *ListNode {
	nd1, nd2 := &ListNode{Val: 0}, &ListNode{Val: 0}
	node1, node2, temp, flag := nd1, nd2, head, 0
	for temp != nil {
		if flag%2 == 0 {
			node1.Next = temp
			node1 = node1.Next
		} else {
			node2.Next = temp
			node2 = node2.Next
		}
		temp = temp.Next
		flag++
	}
	node2.Next = nil
	node1.Next = nd2.Next
	return nd1.Next
}
