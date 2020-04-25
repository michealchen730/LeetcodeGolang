package main

func findMinFibonacciNumbers(k int) int {
	arr := []int{1, 1}
	n := arr[len(arr)-1]
	for n < k {
		n = arr[len(arr)-1] + arr[len(arr)-2]
		arr = append(arr, n)
	}
	res := 0
	for k != 0 {
		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i] <= k {
				k -= arr[i]
				res++
			}
		}
	}
	return res
}
