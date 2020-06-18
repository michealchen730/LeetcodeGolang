package main

func detectCycle(head *ListNode) *ListNode {
	//指定快慢指针
	fast, slow := head, head
	//如果无环，fast会先到达链表的尾部
	for fast != nil {
		//两指针各走一步
		fast, slow = fast.Next, slow.Next
		//如果fast不为空，fast再走一步
		if fast != nil {
			fast = fast.Next
		}
		//实现指针快慢后，如果fast==slow，那么可以确认是环了
		if fast == slow {
			break
		}
	}
	if fast == nil {
		return fast
	}
	//此时重新让fast变成head，注意，这里已经没有快慢这一说法了，纯粹了为了简洁才继续使用fast这一变量
	fast = head
	//两指针各走一步，在k步后会同时到达环的起点，这部分推导还是题解区找下答案吧
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
