package main

func reverseKGroup(head *ListNode, k int) *ListNode {

	if k <= 1 {
		return head
	}
	node := head
	var arr []*ListNode
	for i := 0; i < k; i++ {
		if node != nil {
			arr = append(arr, node)
			node = node.Next
		} else {
			return head
		}
	}
	if len(arr) == k {
		for i := k - 1; i > 0; i-- {
			arr[i].Next = arr[i-1]
		}
	}
	arr[0].Next = reverseKGroup(node, k)
	return arr[k-1]
}
