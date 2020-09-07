package main

//与面试题0207一致

//这题当然还可以用哈希表，但不符合空间复杂度O(1)
func getIntersectionNode160(headA, headB *ListNode) *ListNode {
	t1, t2 := headA, headB
	f1, f2 := true, true
	//简单来说：
	//当headA走完自己的路程后，让它继续走headB的路程
	//当headB走完自己的路程后，让它继续走headA的路程
	//以上两步同时发生，如果存在相交的情况，那么最终会走到一起

	//因为样例中存在链表不相交的情况
	//所以上面两步只允许发生一次
	for t1 != nil && t2 != nil {
		if t1 == t2 {
			return t1
		}
		t1 = t1.Next
		t2 = t2.Next
		if t1 == nil && f1 {
			f1 = false
			t1 = headB
		}
		if t2 == nil && f2 {
			f2 = false
			t2 = headA
		}
	}
	return nil
}
