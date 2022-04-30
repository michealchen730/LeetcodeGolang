package main

//看了下我的java实现，是先转换为二进制，计算1的数目，然后符合条件就加一
//思路估计是看了题解才做出来的
//正确解法好像也就是暴力
func countPrimeSetBits(left int, right int) int {
	res := 0
	for t := left; t <= right; t++ {
		if 1<<onesCount(t)&665772 != 0 {
			res++
		}
	}
	return res
}

//不考虑i为负数的情况
func onesCount(i int) int {
	res := 0
	for i > 0 {
		if i&1 != 0 {
			res++
		}
		i >>= 1
	}
	return res
}
