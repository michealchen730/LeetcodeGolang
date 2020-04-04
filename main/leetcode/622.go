package main

type MyCircularQueue struct {
	stack  []int
	length int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
//func Constructor(k int) MyCircularQueue {
//	return MyCircularQueue{
//		stack:  []int{},
//		length: k,
//	}
//}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.stack = append(this.stack, value)
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.stack = this.stack[1:]
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.stack[0]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.stack) == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return len(this.stack) == this.length
}
