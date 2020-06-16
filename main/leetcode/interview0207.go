package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	t1, t2 := headA, headB
	flag1, flag2 := true, true
	for t1 != nil && t2 != nil {
		if t1 == t2 {
			return t1
		}
		t1 = t1.Next
		t2 = t2.Next
		if t1 == nil && flag1 {
			t1 = headB
			flag1 = false
		}
		if t2 == nil && flag2 {
			t2 = headA
			flag2 = false
		}
	}
	return nil
}
