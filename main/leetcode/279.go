package main

import "math"

//这题没写过题解，导致我竟然看不懂
//看懂了，这题的思路是借鉴BFS的思想，很巧妙，这是我的思路吗？
//将尺寸为n的背包拆解为多个更小的尺寸不同的背包，每一轮会拆多个，level记录当前是第几层
//func numSquares(n int) int {
//	if n == 0 {
//		return 0
//	}
//	arr := []int{n}
//	level := 0
//	for len(arr) != 0 {
//		level++
//		length := len(arr)
//		for i := 0; i < length; i++ {
//			temp := arr[i]
//			//开根号？
//			j := int(math.Sqrt(float64(temp)))
//			for j >= 1 {
//				if j*j == temp {
//					return level
//				}
//				//这里有个拼接操作？
//				arr = append(arr, temp-j*j)
//				j--
//			}
//		}
//		arr = arr[length:]
//	}
//	return level
//}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for k, _ := range dp {
		dp[k] = k
	}
	ltd := int(math.Sqrt(float64(n)))
	for i := 1; i <= ltd; i++ {
		for t := i * i; t <= n; t++ {
			dp[t] = min(dp[t], dp[t-i*i]+1)
		}
	}
	return dp[n]
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
