package main

func findMagicIndex(nums []int) int {
	for k, v := range nums {
		if k == v {
			return k
		}
	}
	return -1
}
