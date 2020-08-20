package main

func smallestK(arr []int, k int) []int {
	var res []int
	if k == 0 {
		return res
	}
	if k >= len(arr) {
		return arr
	}
	for i := 0; i < k; i++ {
		res = append(res, arr[i])
	}
	buildMaxHeap(res)
	for i := k; i < len(arr); i++ {
		if arr[i] < res[0] {
			res[0] = arr[i]
			adjustHeap(res, 0, k)
		}
	}
	return res
}
