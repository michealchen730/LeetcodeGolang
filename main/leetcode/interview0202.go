package main

//上来就简单粗暴的来个暴力法
//这题应该想考察双指针
func kthToLast(head *ListNode, k int) int {
	var nodes []int
	for head != nil {
		nodes = append(nodes, head.Val)
		head = head.Next
	}
	return nodes[len(nodes)-k+1]
}

func kthToLast2(head *ListNode, k int) int {
	t := 0
	temp := head
	for temp != nil {
		temp = temp.Next
		t++
		if t >= k {
			head = head.Next
		}
	}
	return head.Val
}
