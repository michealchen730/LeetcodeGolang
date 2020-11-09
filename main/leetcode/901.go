package main

type StockSpanner struct {
	days  int
	stack []int
	arr   []int
}

//func Constructor() StockSpanner {
//	return StockSpanner{
//		days: 0,
//		stack: make([]int,0),
//		arr: make([]int,0),
//	}
//}

func (this *StockSpanner) Next(price int) int {
	this.days++
	this.arr = append(this.arr, price)
	for len(this.stack) > 0 && this.arr[this.stack[len(this.stack)-1]] <= price {
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, this.days-1)
	if len(this.stack) >= 2 {
		return this.stack[len(this.stack)-1] - this.stack[len(this.stack)-2]
	} else {
		return this.stack[len(this.stack)-1] + 1
	}
}
