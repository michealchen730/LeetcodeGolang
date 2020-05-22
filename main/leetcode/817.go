package main

func numComponents(head *ListNode, G []int) int {
	mp := make(map[int]int)
	for _, v := range G {
		mp[v]++
	}
	res := 0
	flag := 0
	for head != nil {
		if mp[head.Val] > 0 {
			if flag == 0 {
				res++
				flag = 1
			}
		} else {
			flag = 0
		}
		head = head.Next
	}
	return res
}
