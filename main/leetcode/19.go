package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp:=head
	var node *ListNode
	i:=0
	for temp.Next!=nil {
		temp=temp.Next
		i++
		if node!=nil {
			node=node.Next
		}
		if n==i {
			node=head
		}
	}
	if node==nil {
		return head.Next
	}
	node.Next=node.Next.Next
	return head
}
