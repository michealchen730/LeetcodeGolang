package main

func integerBreak(n int) int {
	if n <= 2 {
		return 1
	}
	if n == 3 {
		return 3
	}
	arr := make([]int, n+1)
	//从5开始，所有的乘数i都可以被分解了
	for i := 0; i <= 4; i++ {
		arr[i] = i
	}
	for i := 5; i <= n; i++ {
		temp := i
		for j := 2; j <= i/2; j++ {
			if arr[j]*arr[i-j] > temp {
				temp = arr[j] * arr[i-j]
			}
		}
		arr[i] = temp
	}
	return arr[n]
}
