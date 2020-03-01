package main

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp1 := head
	res := head
	temp2 := head.Next
	temp1.Next = nil
	for temp2 != nil {
		//预留下一个点
		tNext := temp2.Next
		temp2.Next = nil
		var p1 *ListNode
		p2 := res
		for p2 != nil && temp2.Val > p2.Val {
			p1 = p2
			p2 = p2.Next
		}
		temp2.Next = p2
		if p1 == nil {
			res = temp2
		} else {
			p1.Next = temp2
		}
		temp2 = tNext
	}
	return res
}
