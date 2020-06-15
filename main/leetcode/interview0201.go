package main

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head != nil {
		mp := make(map[int]int)
		temp := head
		mp[head.Val]++
		for temp.Next != nil {
			if mp[temp.Next.Val] == 0 {
				mp[temp.Next.Val]++
				temp = temp.Next
			} else {
				temp.Next = temp.Next.Next
			}
		}
	}
	return head
}
