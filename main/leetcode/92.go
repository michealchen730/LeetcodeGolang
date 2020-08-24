package main

//这里给出题解的第二种方法
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//找到位置
	var prev, con, tail *ListNode
	cur := head
	idx_cur := 1
	for idx_cur <= n {
		if idx_cur == m {
			con = prev
			tail = cur
		}
		third := cur.Next
		if idx_cur >= m {
			cur.Next = prev
		}
		idx_cur++
		prev = cur
		cur = third
	}
	tail.Next = cur
	//如果从头开始反转链表，那么需要特判
	if con == nil {
		return prev
	}

	con.Next = prev
	return head
}
