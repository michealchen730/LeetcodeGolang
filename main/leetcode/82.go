package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, flag := head, 0
	var p2, node *ListNode
	for p1 != nil {
		if p1.Next != nil && p1.Next.Val == p1.Val {
			temp, p3 := p1.Val, p1.Next.Next
			for p3 != nil {
				if p3.Val != temp {
					break
				}
				p3 = p3.Next
			}
			if p2 != nil {
				p2.Next = p3
			} else {
				node, flag = p3, 1
			}
			p1 = p3
			continue
		}
		p2 = p1
		p1 = p1.Next
	}
	if node == nil && flag == 0 {
		return head
	}
	return node
}
