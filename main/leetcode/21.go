package main

import "fmt"

//输出单链表的内容
func readLinkList(head *ListNode) {
	for {
		if head == nil {
			break
		}
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{Val: 0}
	head := res
	for {
		if l1 == nil {
			res.Next = l2
			break
		}
		if l2 == nil {
			res.Next = l1
			break
		}
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				res.Next = l1
				l1 = l1.Next
				res = res.Next
			} else {
				res.Next = l2
				l2 = l2.Next
				res = res.Next
			}
		}
	}
	return head.Next
}
