package main

//取巧的做法，显然不符合题意
//func sortList(head *ListNode) *ListNode{
//	var arr []int
//	node:=head
//	for node!=nil{
//		arr=append(arr,node.Val)
//		node=node.Next
//	}
//	sort.Ints(arr)
//	node=head
//	for _,v:=range arr{
//		node.Val=v
//		node=node.Next
//	}
//	return head
//}

func sortList(head *ListNode) *ListNode {
	length := 0
	l := head
	for l != nil {
		length++
		l = l.Next
	}
	dummyHead := &ListNode{Val: 0}
	dummyHead.Next = head
	for s := 1; s < length; s *= 2 {
		cur := dummyHead.Next
		tail := dummyHead
		for cur != nil {
			left := cur
			right := cutNofList(left, s)
			cur = cutNofList(right, s)
			tail.Next = mergeTwoLists(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return dummyHead.Next
}

func cutNofList(l *ListNode, n int) *ListNode {
	head, n := l, n-1
	for head != nil && n != 0 {
		n--
		head = head.Next
	}
	if head == nil {
		return head
	}
	next := head.Next
	head.Next = nil
	return next
}
