package main

//type MyQueue struct {
//	inStack,outStack []int
//}
//
//
///** Initialize your data structure here. */
//func Constructor() MyQueue {
//	return MyQueue{}
//}
//
//
///** Push element x to the back of queue. */
//func (this *MyQueue) Push(x int)  {
//	this.inStack=append(this.inStack,x)
//}
//
//
///** Removes the element from in front of queue and returns that element. */
//func (this *MyQueue) Pop() int {
//	if len(this.outStack)==0{
//		this.in2out()
//	}
//	res:=this.outStack[len(this.outStack)-1]
//	this.outStack=this.outStack[:len(this.outStack)-1]
//	return res
//}
//
//func (this *MyQueue) in2out(){
//	for len(this.inStack)!=0{
//		this.outStack=append(this.outStack,this.inStack[len(this.inStack)-1])
//		this.inStack=this.inStack[:len(this.inStack)-1]
//	}
//}
//
//
///** Get the front element. */
//func (this *MyQueue) Peek() int {
//	if len(this.outStack)==0{
//		this.in2out()
//	}
//	return this.outStack[len(this.outStack)-1]
//}
//
//
///** Returns whether the queue is empty. */
//func (this *MyQueue) Empty() bool {
//	return len(this.inStack)==0&&len(this.outStack)==0
//}
