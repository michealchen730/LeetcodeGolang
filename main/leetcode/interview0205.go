package main

func addTwoNumbers0205(l1 *ListNode, l2 *ListNode) *ListNode {
	flag, res := 0, &ListNode{Val: 0}
	temp := res
	for l1 != nil && l2 != nil {
		s := l1.Val + l2.Val + flag
		if s >= 10 {
			s, flag = s%10, 1
		} else {
			flag = 0
		}
		l1, l2, temp.Next = l1.Next, l2.Next, &ListNode{Val: s}
		temp = temp.Next
	}
	settle := func(node *ListNode) {
		for node != nil {
			s := node.Val + flag
			if s >= 10 {
				s, flag = s%10, 1
			} else {
				flag = 0
			}
			node, temp.Next = node.Next, &ListNode{Val: s}
			temp = temp.Next
		}
	}
	settle(l1)
	settle(l2)
	if flag == 1 {
		temp.Next = &ListNode{Val: 1}
	}
	return res.Next
}
