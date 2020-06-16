package main

func partition0204(head *ListNode, x int) *ListNode {
	node1 := &ListNode{Val: 0}
	node2 := &ListNode{Val: 0}
	t1, t2 := node1, node2
	for head != nil {
		if head.Val < x {
			t1.Next = &ListNode{Val: head.Val}
			t1 = t1.Next
		} else {
			t2.Next = &ListNode{Val: head.Val}
			t2 = t2.Next
		}
		head = head.Next
	}
	t1.Next = node2.Next
	return node1.Next
}
