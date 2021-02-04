package main

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	temp := 0
	for k, v := range nums {
		if temp == sum-v-temp && k != 0 {
			return k
		}
		temp += v
	}
	return -1
}
