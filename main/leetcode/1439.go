package main

import "sort"

func kthSmallest(mat [][]int, k int) int {
	arr1 := make([]int, len(mat[0]))
	copy(arr1, mat[0])
	for i := 1; i < len(mat); i++ {
		flag := 0
		var arr2 []int
		for _, v := range mat[i] {
			for j := 0; j < len(arr1); j++ {
				if len(arr2) < 200 {
					arr2 = append(arr2, v+arr1[j])
				}
				if len(arr2) == 200 {
					if flag == 0 {
						buildMaxHeap(arr2)
						flag = 1
					}
					if v+arr1[j] < arr2[0] {
						arr2[0] = v + arr1[j]
						adjustMaxHeap(arr2, 0, 200)
					} else {
						break
					}
				}
			}
		}
		arr1 = arr2
		sort.Ints(arr1)
	}
	sort.Ints(arr1)
	return arr1[k-1]
}
func adjustMaxHeap(arr []int, i int, length int) {
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && arr[k] < arr[k+1] {
			k = k + 1
		}
		if arr[k] > arr[i] {
			arr[k], arr[i] = arr[i], arr[k]
			i = k
		} else {
			break
		}
	}
}
