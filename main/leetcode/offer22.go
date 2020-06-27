package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	t := 0
	temp := head
	for temp != nil {
		temp = temp.Next
		t++
		if t > k {
			head = head.Next
		}
	}
	return head
}
