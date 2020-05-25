package main

import "fmt"

type LRUCache struct {
	cap int
	mp  map[int]*Node146
	dl  *DoubleList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{mp: make(map[int]*Node146), cap: capacity, dl: CreateDL()}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.mp[key]; ok {
		this.dl.AddLast(v)
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.mp[key]; ok {
		node.val = value
		this.dl.AddLast(node)
	} else {
		node := &Node146{val: value, key: key}
		this.mp[key] = node
		this.dl.tail.prev.next = node
		node.prev = this.dl.tail.prev
		node.next = this.dl.tail
		this.dl.tail.prev = node
		if len(this.mp) > this.cap {
			delete(this.mp, this.dl.RemoveFirst().key)
		}
	}
}

type DoubleList struct {
	head, tail *Node146
}

type Node146 struct {
	prev, next *Node146
	key, val   int
}

func CreateDL() *DoubleList {
	head, tail := &Node146{}, &Node146{}
	head.next, tail.prev = tail, head
	return &DoubleList{
		head: head,
		tail: tail,
	}
}

func (this *DoubleList) AddLast(node *Node146) {
	//node从原来的位置移除
	this.Remove(node)
	//node加入到last
	this.tail.prev.next = node
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev = node
}

func (this *DoubleList) Remove(node *Node146) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}

func (this *DoubleList) RemoveFirst() *Node146 {
	if this.IsEmpty() {
		return nil
	}
	first := this.head.next
	this.Remove(first)
	return first
}

func (this *DoubleList) IsEmpty() bool {
	return this.head.next == this.tail
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回  3
	fmt.Println(cache.Get(4))

}
