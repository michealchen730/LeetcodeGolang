package main

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	fir := head
	for fir != nil {
		sec := fir.Next
		if sec == nil {
			break
		}
		if sec.Val == val {
			fir.Next = sec.Next
			sec.Next = nil
		} else {
			fir = fir.Next
		}
	}
	return head
}
