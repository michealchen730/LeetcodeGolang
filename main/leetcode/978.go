package main

func maxTurbulenceSize(arr []int) int {
	res, temp := 1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			res = max(res, temp)
			temp = 1
			continue
		}
		if i == 1 || (arr[i] > arr[i-1] && arr[i-1] < arr[i-2]) || (arr[i] < arr[i-1] && arr[i-1] > arr[i-2]) {
			temp++
			res = max(res, temp)
			continue
		}
		temp = 2
	}
	return res
}
