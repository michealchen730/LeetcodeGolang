package main

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0
	for k, v := range nums {
		if k%2 == 0 {
			res += v
		}
	}
	return res
}
