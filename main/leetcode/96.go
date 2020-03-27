package main

func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	arr := make([]int, n+1)
	arr[0], arr[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for j := 0; j < i; j++ {
			arr[i] += arr[j] * arr[i-j-1]
		}
	}
	return arr[n]
}
