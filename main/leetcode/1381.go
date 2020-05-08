package main

type CustomStack struct {
	stack   []int
	maxSize int
}

func Constructor1381(maxSize int) CustomStack {
	return CustomStack{
		stack:   []int{},
		maxSize: maxSize,
	}

}

func (this *CustomStack) Push(x int) {
	if len(this.stack) < this.maxSize {
		this.stack = append(this.stack, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(this.stack) == 0 {
		return -1
	} else {
		temp := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		return temp
	}
}

func (this *CustomStack) Increment(k int, val int) {
	if len(this.stack) < k {
		k = len(this.stack)
	}
	for i := 0; i < k; i++ {
		this.stack[i] += val
	}
}
