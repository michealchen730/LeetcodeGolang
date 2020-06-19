package main

//应该是考察单调栈
type MinStack0302 struct {
	stack []int
	mins  []int
}

/** initialize your data structure here. */
func Constructor0302() MinStack0302 {
	return MinStack0302{
		stack: []int{},
		mins:  []int{}, //存的是元素的数组下标
	}
}

func (this *MinStack0302) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.mins) == 0 || this.stack[this.mins[len(this.mins)-1]] > x {
		this.mins = append(this.mins, len(this.stack)-1)
	}
}

func (this *MinStack0302) Pop() {
	if len(this.stack) != 0 {
		this.stack = this.stack[:len(this.stack)-1]
		for len(this.mins) > 0 && this.mins[len(this.mins)-1] >= len(this.stack) {
			this.mins = this.mins[:len(this.mins)-1]
		}
	}
}

func (this *MinStack0302) Top0302() int {
	if len(this.stack) != 0 {
		return this.stack[len(this.stack)-1]
	}
	return -1
}

func (this *MinStack0302) GetMin0302() int {
	if len(this.mins) != 0 {
		return this.stack[this.mins[len(this.mins)-1]]
	}
	return -1
}
