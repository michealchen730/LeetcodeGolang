package main

type MyQueue struct {
	queue []int
}

/** Initialize your data structure here. */
func ConstructorI0304() MyQueue {
	return MyQueue{queue: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.queue) != 0 {
		temp := this.queue[0]
		this.queue = this.queue[1:]
		return temp
	}
	return -1
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.queue) != 0 {
		return this.queue[0]
	}
	return -1
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}
