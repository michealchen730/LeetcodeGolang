package main

type RLEIterator struct {
	cnt  []int
	arr  []int
	temp int
}

func Constructor900(A []int) RLEIterator {
	length := len(A) / 2
	c, a := make([]int, length), make([]int, length)

	for k := 0; k < length; k++ {
		c[k] = A[k*2]
		a[k] = A[k*2+1]
	}

	return RLEIterator{
		cnt:  c,
		arr:  a,
		temp: 0,
	}
}

func (this *RLEIterator) Next(n int) int {
	for this.temp < len(this.cnt) && this.cnt[this.temp] < n {
		n -= this.cnt[this.temp]
		this.temp++
	}

	if this.temp == len(this.cnt) {
		return -1
	}

	this.cnt[this.temp] -= n

	return this.arr[this.temp]
}
