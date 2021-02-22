package main

func findMaxConsecutiveOnes(nums []int) int {
	res, tmp := 0, 0
	for _, v := range nums {
		if v == 0 {
			tmp = 0
		} else {
			tmp += 1
			res = max(res, tmp)
		}
	}
	return res
}
