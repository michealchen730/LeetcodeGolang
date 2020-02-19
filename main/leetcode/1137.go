package main

func tribonacci(n int) int {
	res := make([]int, 3)
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	res[1], res[2] = 1, 1
	for i := 3; i <= n; i++ {
		res = append(res, res[i-1]+res[i-2]+res[i-3])
	}
	return res[n]
}
