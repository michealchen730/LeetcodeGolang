package main

func rotateRight(head *ListNode, k int) *ListNode {
	if k==0||head==nil||head.Next==nil {
		return head
	}
	length,temp:=0,head
	for temp.Next!=nil{
		temp=temp.Next
		length++
	}
	length++
	var res *ListNode
	last,i:=temp,0
	temp=head
	k=k%length
	if k==0 {
		return head
	}else {
		last.Next=head
		k=length-k-1
		for temp!=nil {
			if i==k {
				res=temp.Next
				temp.Next=nil
				return res
			}else {
				temp=temp.Next
				i++
			}
		}
	}
	return res
}
