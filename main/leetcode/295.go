package main

type MedianFinder struct {
	maxh MaxHeap
	minh MinHeap
}

/** initialize your data structure here. */
//func Constructor() MedianFinder {
//	return MedianFinder{
//		maxh: MaxHeap{heap: []int{},size: 0},
//		minh: MinHeap{heap: []int{},size: 0},
//	}
//}

func (this *MedianFinder) AddNum(num int) {
	this.maxh.push(num)

	this.minh.push(this.maxh.pop())

	if this.maxh.size < this.minh.size {
		this.maxh.push(this.minh.pop())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxh.size > this.minh.size {
		return float64(this.maxh.heap[0])
	} else {
		return (float64(this.maxh.heap[0]) + float64(this.minh.heap[0])) / 2
	}
}
