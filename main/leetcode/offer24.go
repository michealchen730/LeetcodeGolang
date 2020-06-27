package main

//递归
func reverseListO24(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	temp := head
	if temp.Next == nil {
		return temp
	}
	res := reverseListO24(temp.Next)
	temp.Next.Next = temp
	temp.Next = nil
	return res
}
