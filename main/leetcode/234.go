package main

//快慢指针定位中间节点
//反转后半部分链表
//判定
//还原后半部分链表
func isPalindrome234(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	backup := slow

	//得到反转后的尾节点
	rear := reverseList(backup)
	h1, h2 := head, rear
	for h2 != nil {
		if h1.Val != h2.Val {
			return false
		} else {
			h1 = h1.Next
			h2 = h2.Next
		}
	}
	reverseList(rear)
	return true
}
