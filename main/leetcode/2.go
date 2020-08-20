package main

//最早的写法
//逐步相加，变成数组，然后利用数组直接构建结点
func addTwoNumbers02(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum []int
	var temp = 0
	var flag = 0
	for {
		if l1 == nil || l2 == nil {
			break
		} else {
			temp = flag + l1.Val + l2.Val
			if temp >= 10 {
				flag = 1
			} else {
				flag = 0
			}
			sum = append(sum, temp%10)
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	if l1 == nil {
		second(l2, &flag, &temp, &sum)
	}
	if l2 == nil {
		second(l1, &flag, &temp, &sum)
	}
	if flag == 1 {
		sum = append(sum, 1)
	}
	return buildLinkList(sum)
}

func second(l1 *ListNode, flag *int, temp *int, sum *[]int) {
	for {
		if l1 == nil {
			break
		} else {
			*temp = l1.Val + *flag
			if *temp >= 10 {
				*flag = 1
			} else {
				*flag = 0
			}
			*sum = append(*sum, *temp%10)
			l1 = l1.Next
		}
	}
}

//还是重建结点的方式
func addTwoNumbers03(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyhead := &ListNode{Val: 0}
	carry, node := 0, dummyhead
	for l1 != nil || l2 != nil {
		x1, x2 := 0, 0
		if l1 != nil {
			x1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			x2 = l2.Val
			l2 = l2.Next
		}
		sum := x1 + x2 + carry
		carry = sum / 10
		sum = sum % 10
		node.Next = &ListNode{Val: sum}
		node = node.Next
	}
	//注意最后是否还存在进位
	if carry != 0 {
		node.Next = &ListNode{Val: carry}
	}
	return dummyhead.Next
}
