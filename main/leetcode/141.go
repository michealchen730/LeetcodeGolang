package main

func hasCycle(head *ListNode) bool {
	if head==nil||head.Next==nil{
		return false
	}
	pointer1,pointer2:=head,head
	for {
		if pointer2.Next!=nil&&pointer2.Next.Next!=nil {
			pointer2=pointer2.Next.Next
			pointer1=pointer1.Next
			if pointer1==pointer2 {
				return true
			}
		}else {
			return false
		}
	}
}