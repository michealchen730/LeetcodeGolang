package main

func deleteNode237(node *ListNode) {
	pre := node
	for node.Next != nil {
		node.Val = node.Next.Val
		pre = node
		node = node.Next
	}
	pre.Next = nil
}
