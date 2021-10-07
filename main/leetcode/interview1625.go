package main

//type DoublyLinkedList struct {
//	key int
//	val int
//	prv *DoublyLinkedList
//	nxt *DoublyLinkedList
//}
//
//type LRUCache struct {
//	mp map[int]*DoublyLinkedList//确定链表是否存在的
//	head *DoublyLinkedList
//	tail *DoublyLinkedList
//	capacity int
//}
//
//
//func Constructor(capacity int) LRUCache {
//	mp:=make(map[int]*DoublyLinkedList)
//	head:=DoublyLinkedList{
//		key: -1,
//		val: -1,
//	}
//	tail:=DoublyLinkedList{
//		key: -1,
//		val: -1,
//		prv: &head,
//	}
//	head.nxt=&tail
//	return LRUCache{
//		mp: mp,
//		head: &head,
//		tail: &tail,
//		capacity: capacity,
//	}
//}
//
//
//func (this *LRUCache) Get(key int) int {
//	if node,ok:=this.mp[key];ok{
//		//假设放在尾节点
//		//首先把当前节点剥离出来
//		node.prv.nxt=node.nxt
//		node.nxt.prv=node.prv
//		//再放到尾部去
//		this.tail.prv.nxt=node
//		node.prv=this.tail.prv
//		node.nxt=this.tail
//		this.tail.prv=node
//		return node.val
//	}else{
//		return -1
//	}
//}
//
//
//func (this *LRUCache) Put(key int, value int)  {
//	var node *DoublyLinkedList
//	if _,ok:=this.mp[key];!ok{
//		//容量满了，就清除掉最近未使用的元素
//		if len(this.mp)==this.capacity{
//			delete(this.mp,this.head.nxt.key)
//			this.head.nxt.nxt.prv=this.head
//			this.head.nxt=this.head.nxt.nxt
//		}
//		node=&DoublyLinkedList{
//			key: key,
//			val: value,
//			prv: this.tail.prv,
//			nxt: this.tail,
//		}
//		this.mp[key]=node
//	}else{
//		node=this.mp[key]
//		node.val=value
//		//把node剥离出来
//		node.prv.nxt=node.nxt
//		node.nxt.prv=node.prv
//	}
//	//把node放到尾部去
//	this.tail.prv.nxt=node
//	node.prv=this.tail.prv
//	node.nxt=this.tail
//	this.tail.prv=node
//}
