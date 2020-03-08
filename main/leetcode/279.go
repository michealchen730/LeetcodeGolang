package main

import "math"

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	arr := []int{n}
	level := 0
	for len(arr) != 0 {
		level++
		length := len(arr)
		for i := 0; i < length; i++ {
			temp := arr[i]
			j := int(math.Sqrt(float64(temp)))
			for j >= 1 {
				if j*j == temp {
					return level
				}
				arr = append(arr, temp-j*j)
				j--
			}
		}
		arr = arr[length:]
	}
	return level
}

//func numSquares(n int) int {
//	dp:=make([]int,n+1)
//	dp[1]=1
//	for i:=2; i<=n; i++ {
//		dp[i]=i
//		for j:=1;j*j<=i;j++{
//			dp[i]=min(dp[i],dp[i-j*j]+1)
//		}
//	}
//	return dp[n]
//}
