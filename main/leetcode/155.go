package main

type MinStack struct {
	stack    []int
	minstack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var stack, minstack []int
	return MinStack{
		stack:    stack,
		minstack: minstack,
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minstack) != 0 && x > this.minstack[len(this.minstack)-1] {
		x = this.minstack[len(this.minstack)-1]
	}
	this.minstack = append(this.minstack, x)
}

func (this *MinStack) Pop() {
	length := len(this.stack)
	this.stack = this.stack[:length-1]
	this.minstack = this.minstack[:length-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}
