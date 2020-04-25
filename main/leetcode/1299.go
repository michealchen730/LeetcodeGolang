package main

func replaceElements(arr []int) []int {
	mx := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		temp := arr[i]
		arr[i] = mx
		mx = max(temp, mx)
	}
	return arr
}
