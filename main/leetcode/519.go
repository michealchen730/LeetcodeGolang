package main

import (
	"math/rand"
)

type Solution519 struct {
	mp                  map[int]int
	n_rows, n_cols, rem int
}

func Constructor519(n_rows int, n_cols int) Solution519 {
	//要把二维矩阵映射到一维
	return Solution519{
		mp:     make(map[int]int),
		n_rows: n_rows,
		n_cols: n_cols,
		rem:    n_rows * n_cols,
	}
}

func (this *Solution519) Flip() []int {
	//每次生成一个值t
	t := rand.Intn(this.rem)
	//这个值需要做--操作
	this.rem--
	var x int
	//如果mp[t]存在，即这个值已经为1了
	if v, ok := this.mp[t]; ok {
		x = v
	} else {
		//不存在，则x置t
		x = t
	}
	var val int
	if v, ok := this.mp[this.rem]; ok {
		val = v
	} else {
		val = this.rem
	}
	this.mp[t] = val
	return []int{x / this.n_cols, x % this.n_cols}
}

func (this *Solution519) Reset() {
	this.rem = this.n_cols * this.n_rows
	this.mp = map[int]int{}
}
