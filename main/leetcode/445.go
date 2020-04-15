package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList2(l1)
	l2 = reverseList2(l2)
	temp := addTwoLinkedLists(l1, l2)
	return reverseList2(temp)
}

//两个链表相加，使用的某大佬的代码
func addTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var temp = &ListNode{}
	var incr int

	lastNode := temp
	for {
		var v int
		if l1 != nil && l2 != nil {
			v = l1.Val + l2.Val + incr
			if v > 9 {
				v = v - 10
				incr = 1
			} else {
				incr = 0
			}
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			v = l1.Val + incr
			if v > 9 {
				v = v - 10
				incr = 1
			} else {
				incr = 0
			}
			l1 = l1.Next
		} else if l2 != nil {
			v = l2.Val + incr
			if v > 9 {
				v = v - 10
				incr = 1
			} else {
				incr = 0
			}
			l2 = l2.Next
		} else {
			break
		}
		node := &ListNode{
			Val: v,
		}
		lastNode.Next = node

		lastNode = node
	}

	if incr == 1 {
		node := &ListNode{
			Val: 1,
		}
		lastNode.Next = node
	}

	return temp.Next

}

//翻转链表，用的某位大佬的代码
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode = nil
	var cur *ListNode = head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
