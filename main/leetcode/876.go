package main

func middleNode(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p2 != nil {
		p2 = p2.Next
		if p2 != nil {
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return p1
}
