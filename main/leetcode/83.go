package main

func deleteDuplicates83(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre := head
	temp := head.Next
	for temp != nil {
		if temp.Val == pre.Val {
			pre.Next = temp.Next
		} else {
			pre = pre.Next
		}
		temp = temp.Next
	}
	return head
}
