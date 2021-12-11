package main

//这题我看得懂，但是我不会做
//func kInversePairs(n int, k int) int {
//	dp:=make([][]int,n+1)
//	for i,_:=range dp{
//		dp[i]=make([]int,k+1)
//	}
//	dp[0][0]=1
//	for ida:=1;ida<=n;ida++{
//		for idb:=0;idb<=k;idb++{
//			for t:=1;t<=ida;t++{
//				if idb-ida+t>=0 {
//					dp[ida][idb] += dp[ida-1][idb-ida+t]
//					dp[ida][idb]%=1000000007
//				}
//			}
//		}
//	}
//	return dp[n][k]
//}

//时间优化版
//注意：这里超过1000000007要减去1000000007，小于0则要加上1000000007
//func kInversePairs(n int, k int) int {
//	dp:=make([][]int,n+1)
//	for i,_:=range dp{
//		dp[i]=make([]int,k+1)
//		dp[i][0]=1
//	}
//	for ida:=1;ida<=n;ida++{
//		for idb:=1;idb<=k;idb++{
//			dp[ida][idb] = dp[ida][idb-1]+dp[ida-1][idb]
//			if idb-ida>=0{
//				dp[ida][idb]-=dp[ida-1][idb-ida]
//			}
//			if dp[ida][idb]>=1000000007{
//				dp[ida][idb]-=1000000007
//			}else if dp[ida][idb]<0{
//				dp[ida][idb]+=1000000007
//			}
//		}
//	}
//	return dp[n][k]
//}

func kInversePairs(n int, k int) int {
	dp := make([][]int, 2)
	for i, _ := range dp {
		dp[i] = make([]int, k+1)
		dp[i][0] = 1
	}
	for ida := 1; ida <= n; ida++ {
		for idb := 1; idb <= k; idb++ {
			dp[1][idb] = dp[1][idb-1] + dp[0][idb]
			if idb-ida >= 0 {
				dp[1][idb] -= dp[0][idb-ida]
			}
			if dp[1][idb] >= 1000000007 {
				dp[1][idb] -= 1000000007
			} else if dp[1][idb] < 0 {
				dp[1][idb] += 1000000007
			}
		}
		copy(dp[0], dp[1])
	}
	return dp[1][k]
}
