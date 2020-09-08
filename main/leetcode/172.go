package main

//这个解法太秀了

//通过对n除以5，得到所有5的倍数
//将n变成n除以5
//再重复前两个步骤，就会陆续得到5的幂的倍数

func trailingZeroes(n int) int {
	zeroCount := 0
	for n > 0 {
		n /= 5
		zeroCount += n
	}
	return zeroCount
}
