package main

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	for length != 1 {
		for i := 0; i < length/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[length-i-1])
		}
		if length%2 == 0 {
			length = length / 2
		} else {
			length = length/2 + 1
		}
	}
	return lists[0]
}
