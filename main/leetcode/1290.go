package main

func getDecimalValue(head *ListNode) int {
	sum := 0
	var arr []int
	for head != nil {
		for i := 0; i < len(arr); i++ {
			arr[i] = arr[i] << 1
		}
		arr = append(arr, head.Val)
		head = head.Next
	}
	for _, v := range arr {
		sum += v
	}
	return sum
}
