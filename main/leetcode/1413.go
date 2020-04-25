package main

func minStartValue(nums []int) int {
	sum := 0
	res := 0
	for _, v := range nums {
		sum += v
		if sum < 1 {
			res += 1 - sum
			sum = 1
		}
	}
	if res == 0 {
		res = 1
	}
	return res
}
