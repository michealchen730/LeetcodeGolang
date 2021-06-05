package main

import "fmt"

func heapsort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		AjHeap(arr[:i+1])
		arr[0], arr[i] = arr[i], arr[0]
	}
	fmt.Println(arr)
}

func BuildHeap(arr []int) {
	n := len(arr)
	end := n/2 - 1
	for j := end; j >= 0; j-- {
		//保证右叶子节点存在
		e := j * 2
		if j*2+1 < len(arr) {
			if arr[e] < arr[e+1] {
				e++
			}
		}
		if arr[j] < arr[e] {
			arr[j], arr[e] = arr[e], arr[j]
		}
	}
}

func AjHeap(arr []int) {
	start := 0
	for start*2+1 < len(arr) {
		l := start*2 + 1
		if l+1 < len(arr) {
			if arr[l] < arr[l+1] {
				l++
			}
		}
		if arr[l] > arr[start] {
			arr[start], arr[l] = arr[l], arr[start]
			start = l
		} else {
			break
		}
	}
}
