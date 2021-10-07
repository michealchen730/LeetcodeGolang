package main

import (
	"math"
)

//func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
//	dp:=make([][]int,k+2)
//	for i,_:=range dp{
//		dp[i]=make([]int,n)
//		for j,_:=range dp[i]{
//			dp[i][j]=math.MaxInt32
//		}
//	}
//
//	dp[0][src]=0
//	for c:=1;c<=k+1;c++{
//		for i:=0;i<n;i++{
//			if i==src{
//				continue
//			}
//			for _,v:=range flights{
//				if v[1]==i&&dp[c-1][v[0]]!=math.MaxInt32{
//					dp[c][i]=min(dp[c][i],dp[c-1][v[0]]+v[2])
//				}
//			}
//		}
//	}
//	res:=math.MaxInt32
//	for i:=1;i<=k+1;i++{
//		res=min(res,dp[i][dst])
//	}
//	if res==math.MaxInt32{
//		return -1
//	}
//	return res
//}

//func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
//	dp:=make([][]int,k+2)
//	for i,_:=range dp{
//		dp[i]=make([]int,n)
//		for j,_:=range dp[i]{
//			dp[i][j]=math.MaxInt32
//		}
//	}
//	mp:=make(map[int][][]int)
//	for _,v:=range flights{
//		mp[v[1]]=append(mp[v[1]],[]int{v[0],v[2]})
//	}
//
//	dp[0][src]=0
//	for c:=1;c<=k+1;c++{
//		for i:=0;i<n;i++{
//			if i==src{
//				continue
//			}
//			for _,v:=range mp[i]{
//				if dp[c-1][v[0]]!=math.MaxInt32{
//					dp[c][i]=min(dp[c][i],dp[c-1][v[0]]+v[1])
//				}
//			}
//		}
//	}
//	res:=math.MaxInt32
//	for i:=1;i<=k+1;i++{
//		res=min(res,dp[i][dst])
//	}
//	if res==math.MaxInt32{
//		return -1
//	}
//	return res
//}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dp := make([][]int, 2)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		for j, _ := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	mp := make(map[int][][]int)
	for _, v := range flights {
		mp[v[1]] = append(mp[v[1]], []int{v[0], v[2]})
	}
	res := math.MaxInt32
	dp[0][src] = 0
	for c := 1; c <= k+1; c++ {
		for i := 0; i < n; i++ {
			if i == src {
				continue
			}
			for _, v := range mp[i] {
				if dp[0][v[0]] != math.MaxInt32 {
					dp[1][i] = min(dp[1][i], dp[0][v[0]]+v[1])
				}
			}
		}
		if dp[1][dst] != math.MaxInt32 {
			res = min(res, dp[1][dst])
		}
		for i, v := range dp[1] {
			dp[0][i] = v
			dp[1][i] = math.MaxInt32
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
