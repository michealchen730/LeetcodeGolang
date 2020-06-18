package main

//应该是考察单调栈
type MinStack struct {
	stack []int
	mins  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		mins:  []int{}, //存的是元素的数组下标
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.mins) == 0 || this.stack[this.mins[len(this.mins)-1]] > x {
		this.mins = append(this.mins, len(this.stack)-1)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) != 0 {
		this.stack = this.stack[:len(this.stack)-1]
		for len(this.mins) > 0 && this.mins[len(this.mins)-1] >= len(this.stack) {
			this.mins = this.mins[:len(this.mins)-1]
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) != 0 {
		return this.stack[len(this.stack)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.mins) != 0 {
		return this.stack[this.mins[len(this.mins)-1]]
	}
	return -1
}
