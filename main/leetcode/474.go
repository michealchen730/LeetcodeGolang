package main

//func findMaxForm(strs []string, m int, n int) int {
//	//背包问题
//	dp:=make([][][]int,len(strs)+1)
//	for i,_:=range dp{
//		dp[i]=make([][]int,m+1)
//		for j,_:=range dp[i]{
//			dp[i][j]=make([]int,n+1)
//		}
//	}
//	res:=0
//	for idx:=1;idx<=len(strs);idx++{
//		cm,cn:=countZeroAndOne(strs[idx-1])
//		for m1:=0;m1<=m;m1++{
//			for n1:=0;n1<=n;n1++{
//				if m1>=cm&&n1>=cn{
//					dp[idx][m1][n1]=max(dp[idx-1][m1][n1],dp[idx-1][m1-cm][n1-cn]+1)
//				}else{
//					dp[idx][m1][n1]=dp[idx-1][m1][n1]
//				}
//				res=max(res,dp[idx][m1][n1])
//			}
//		}
//	}
//	return res
//}

//func findMaxForm(strs []string, m int, n int) int {
//	//背包问题
//	dp:=make([][]int,m+1)
//	for i,_:=range dp{
//		dp[i]=make([]int,n+1)
//	}
//	res:=0
//	for idx:=1;idx<=len(strs);idx++{
//		cm,cn:=countZeroAndOne(strs[idx-1])
//		for m1:=m;m1>=0;m1--{
//			for n1:=n;n1>=0;n1--{
//				if m1>=cm&&n1>=cn{
//					dp[m1][n1]=max(dp[m1][n1],dp[m1-cm][n1-cn]+1)
//				}
//				res=max(res,dp[m1][n1])
//			}
//		}
//	}
//	return res
//}

func findMaxForm(strs []string, m int, n int) int {
	//背包问题
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	res := 0
	for idx := 1; idx <= len(strs); idx++ {
		cm, cn := countZeroAndOne(strs[idx-1])
		for m1 := m; m1 >= 0; m1-- {
			for n1 := n; n1 >= 0; n1-- {
				if m1 >= cm && n1 >= cn {
					dp[m1][n1] = max(dp[m1][n1], dp[m1-cm][n1-cn]+1)
				}
				res = max(res, dp[m1][n1])
			}
		}
	}
	return res
}

func countZeroAndOne(str string) (int, int) {
	zero, one := 0, 0
	for _, v := range str {
		if v == '0' {
			zero++
		}
		if v == '1' {
			one++
		}
	}
	return zero, one
}
