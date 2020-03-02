package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//将切片转化为单链表并返回链表头部
func buildLinkList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head = &ListNode{Val: arr[0]}
	var temp = head
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{Val: arr[i]}
		temp = temp.Next
	}
	return head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ht := make([]*ListNode, 2)
	rList(head, ht)
	return ht[0]
}

func rList(head *ListNode, ht []*ListNode) {
	if head.Next == nil {
		ht[0] = head
		ht[1] = head
		return
	}
	rList(head.Next, ht)
	ht[1].Next = head
	head.Next = nil
	ht[1] = head
}
