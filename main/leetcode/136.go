package main

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res = res ^ v
	}
	return res
}
