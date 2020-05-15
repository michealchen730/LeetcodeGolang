package main

func nextLargerNodes(head *ListNode) []int {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	var stack []int
	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, arr[i])
		if len(stack) == 1 {
			arr[i] = 0
		} else {
			arr[i] = stack[len(stack)-2]
		}
	}
	return arr
}
