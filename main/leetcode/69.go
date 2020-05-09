package main

import "math"

//先尝试下二分查找
func mySqrt(x int) int {
	low, high := 0, x
	for low <= high-1 {
		mid1 := low + (high-low)/2
		mid2 := mid1 + 1
		prd1 := mid1 * mid1
		prd2 := mid2 * mid2
		if prd2 == x {
			return mid2
		}
		if prd1 <= x && prd2 > x {
			return mid1
		}
		if prd2 < x {
			low = mid2
		}
		if prd1 > x {
			high = mid1
		}
	}
	return low
}

//牛顿迭代（直接搬运）
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}
