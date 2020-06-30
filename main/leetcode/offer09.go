package main

type CQueue struct {
	stack1 []int
}

func ConstructorO09() CQueue {
	return CQueue{stack1: []int{}}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)

}

func (this *CQueue) DeleteHead() int {
	if len(this.stack1) == 0 {
		return -1
	} else {
		temp := this.stack1[0]
		this.stack1 = this.stack1[1:]
		return temp
	}
}
