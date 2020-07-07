package main

type NodeO35 struct {
	Val    int
	Next   *NodeO35
	Random *NodeO35
}

func copyRandomList(head *NodeO35) *NodeO35 {
	if head == nil {
		return nil
	}
	mp := make(map[*NodeO35]*NodeO35)
	res := head
	var pre *NodeO35
	for head != nil {
		if mp[head] == nil { //如果没有对应元素，那么新建一个
			node := &NodeO35{Val: head.Val}
			//使用map匹配起来
			mp[head] = node
		}
		//与链表的上个元素相连
		if pre != nil {
			mp[pre].Next = mp[head]
		}
		//准备random元素
		if head.Random != nil {
			if mp[head.Random] == nil {
				rand := &NodeO35{Val: head.Random.Val}
				mp[head.Random] = rand
			}
			mp[head].Random = mp[head.Random]
		}
		//用pre记录上一个链表的情况
		pre = head
		head = head.Next
	}
	return mp[res]
}
