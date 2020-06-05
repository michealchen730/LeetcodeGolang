package main

//超时的方法
func rangeBitwiseAnd1(m int, n int) int {
	res := m
	for i := m + 1; i <= n; i++ {
		res &= i
	}
	return res
}

//位运算自身规律
func rangeBitwiseAnd(m int, n int) int {
	shift := 0
	for m < n {
		m >>= 1
		n >>= 1
		shift++
	}
	return m << shift
}

//Brian Kernighan算法
func rangeBitwiseAnd2(m int, n int) int {
	for m < n {
		n = n & (n - 1)
	}
	return m & n
}
