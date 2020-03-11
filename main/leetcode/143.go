package main

func reorderList(head *ListNode) {
	node := head
	if node == nil || node.Next == nil || node.Next.Next == nil {
		return
	}
	var arr []*ListNode
	arr = append(arr, head)
	length := 1
	for node.Next != nil {
		length++
		node = node.Next
		arr = append(arr, node)
	}
	for i := 0; i < length/2; i++ {
		arr[i].Next = arr[len(arr)-1-i]
		arr[len(arr)-1-i].Next = arr[i+1]
	}
	arr[length/2].Next = nil
}
