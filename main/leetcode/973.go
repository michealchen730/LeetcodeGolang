package main

import "math"

func kClosest(points [][]int, K int) [][]int {
	var res [][]int
	var heap []int
	flag := 0
	for _, v := range points {
		temp := int(math.Pow(float64(v[0]), 2)) + int(math.Pow(float64(v[1]), 2))
		if len(res) < K {
			res = append(res, v)
			heap = append(heap, temp)
		} else {
			if flag == 0 {
				BMH973(heap, res)
				flag = 1
			}
			if temp < heap[0] {
				heap[0] = temp
				res[0] = v
				AMH973(heap, 0, K, res)
			}
		}
	}
	return res
}
func BMH973(arr []int, res [][]int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		AMH973(arr, i, len(arr), res)
	}
}
func AMH973(arr []int, i int, length int, res [][]int) {
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && arr[k] < arr[k+1] {
			k = k + 1
		}
		if arr[k] > arr[i] {
			arr[k], arr[i] = arr[i], arr[k]
			res[k], res[i] = res[i], res[k]
			i = k
		} else {
			break
		}
	}
}
