package main

import (
	"math"
)

func nthSuperUglyNumber(n int, primes []int) int {
	arr := make([]int, n)
	arr[0] = 1
	//candi用来记录，每一位的乘primes的当前最小序列
	cand := make([]int, len(primes))
	minArr := make([]int, len(primes))
	//mn表示第idx个位置成功当选

	//每次遍历均更新一个值arr[i]
	for i := 1; i < n; i++ {
		//更新minArr
		tmpMin := math.MaxInt32
		for k, _ := range minArr {
			minArr[k] = primes[k] * arr[cand[k]]
			if minArr[k] < tmpMin {
				tmpMin = minArr[k]
			}
		}
		for k, v := range minArr {
			if v <= tmpMin {
				cand[k]++
			}
		}
		arr[i] = tmpMin
	}
	return arr[n-1]
}
