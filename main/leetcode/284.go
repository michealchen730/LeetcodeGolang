package main

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	Array []int
}

//func Constructor(iter *Iterator) *PeekingIterator {
//	var arr []int
//	for iter.hasNext(){
//		arr=append(arr,iter.next())
//	}
//	return &PeekingIterator{
//		Array: arr,
//	}
//}

func (this *PeekingIterator) hasNext() bool {
	return len(this.Array) > 0
}

func (this *PeekingIterator) next() int {
	res := this.Array[0]
	this.Array = this.Array[1:]
	return res
}

func (this *PeekingIterator) peek() int {
	return this.Array[0]
}
