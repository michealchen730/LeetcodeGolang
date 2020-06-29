package main

func mergeTwoListsO25(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	res := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
		//这里是保险起见
		head.Next = nil
	}
	if l1 == nil {
		head.Next = l2
	} else {
		head.Next = l1
	}
	return res.Next
}
