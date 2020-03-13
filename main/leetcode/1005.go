package main

import (
	"sort"
)

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	temp, k := 0, K
	if A[0] < 0 { //代表有负数，那么先翻转负数部分
		i := 0
		for ; i < K; i++ {
			if A[i] < 0 {
				A[i] = -A[i]
			} else { //如果所有负数都被翻转完了，而K还有剩余，那么停下
				if A[i] < A[i-1] { //此时我们需要知道当前数组的最小值
					temp = i
				} else {
					temp = i - 1
				}
				break
			}
		}
		k = K - i //还需要翻转多少次
	}
	if k%2 != 0 { //到达这一步的时候，要么数组里没有负数，但k还大于等于0；要么数组里还有负数，但K已经被用完了
		A[temp] = -A[temp]
	}
	res := 0
	for _, v := range A {
		res += v
	}
	return res
}
