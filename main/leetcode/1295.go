package main

import "strconv"

func findNumbers(nums []int) int {
	res := 0
	for _, v := range nums {
		s := strconv.Itoa(v)
		if len(s)%2 == 0 {
			res++
		}
	}
	return res
}
