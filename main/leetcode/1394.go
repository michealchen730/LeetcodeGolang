package main

func findLucky(arr []int) int {
	res := -1
	arr2 := make([]int, len(arr)+1)
	for _, v := range arr {
		if v <= len(arr) {
			arr2[v]++
		}
	}
	for k, v := range arr2 {
		if v == k {
			res = v
		}
	}
	if res == 0 {
		return -1
	}
	return res
}
