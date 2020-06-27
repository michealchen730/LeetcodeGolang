package main

func reversePrint(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	reverseArr(res)
	return res
}
