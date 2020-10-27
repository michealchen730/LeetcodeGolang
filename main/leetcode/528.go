package main

import "math/rand"

type Solution528 struct {
	arr []int
	sum int
}

func Constructor528(w []int) Solution528 {
	arr := make([]int, len(w))
	sum := 0
	for i := 0; i < len(w); i++ {
		sum += w[i]
		arr[i] = sum
	}

	return Solution528{
		sum: sum,
		arr: arr,
	}

}

func (this *Solution528) PickIndex() int {
	n := rand.Intn(this.sum) + 1
	for i := 0; i < len(this.arr); i++ {
		if n <= this.arr[i] {
			return i
		}
	}
	return -1
}
