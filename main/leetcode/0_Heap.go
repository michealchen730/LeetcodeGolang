package main

type MaxHeap struct {
	heap []int
	size int
}

func (this *MaxHeap) push(num int) {
	this.heap = append(this.heap, num)
	this.size++
	k := this.size - 1
	for k > 0 {
		child := (k - 1) / 2
		if this.heap[child] < this.heap[k] {
			this.heap[child], this.heap[k] = this.heap[k], this.heap[child]
			k = child
		} else {
			break
		}
	}
}

func (this *MaxHeap) pop() int {
	res := this.heap[0]
	this.heap[0] = this.heap[this.size-1]
	this.heap = this.heap[:this.size-1]
	this.size--
	this.adjustMaxHeap(0)
	return res
}

func (this *MaxHeap) adjustMaxHeap(i int) {
	for k := 2*i + 1; k < this.size; k = 2*k + 1 {
		if k+1 < this.size && this.heap[k] < this.heap[k+1] {
			k = k + 1
		}
		if this.heap[k] > this.heap[i] {
			this.heap[k], this.heap[i] = this.heap[i], this.heap[k]
			i = k
		} else {
			break
		}
	}
}

type MinHeap struct {
	heap []int
	size int
}

func (this *MinHeap) push(num int) {
	this.heap = append(this.heap, num)
	this.size++
	k := this.size - 1
	for k > 0 {
		child := (k - 1) / 2
		if this.heap[child] > this.heap[k] {
			this.heap[child], this.heap[k] = this.heap[k], this.heap[child]
			k = child
		} else {
			break
		}
	}
}

func (this *MinHeap) pop() int {
	res := this.heap[0]
	this.heap[0] = this.heap[this.size-1]
	this.heap = this.heap[:this.size-1]
	this.size--
	this.adjustMinHeap(0)
	return res
}

func (this *MinHeap) adjustMinHeap(i int) {
	for k := 2*i + 1; k < this.size; k = 2*k + 1 {
		if k+1 < this.size && this.heap[k] > this.heap[k+1] {
			k = k + 1
		}
		if this.heap[k] < this.heap[i] {
			this.heap[k], this.heap[i] = this.heap[i], this.heap[k]
			i = k
		} else {
			break
		}
	}
}
