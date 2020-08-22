package main

func minCut(s string) int {
	if len(s) == 0 {
		return 0
	}

	//这题感觉上应该是有个预处理
	//需要看一下所有的字串是否是回文串
	n := len(s)
	predp := make([][]bool, n)
	for k, _ := range predp {
		predp[k] = make([]bool, n)
	}
	//传统的判断算法最终会超时
	//for i:=0;i<n;i++{
	//	for j:=i;j<n;j++{
	//		if isPalindrome(s[i:j+1]){
	//			predp[i][j]=true
	//		}
	//	}
	//}
	for i := 0; i < n; i++ {
		predp[i][i] = true
	}
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				predp[i][j] = false
			} else {
				if j-i < 3 {
					predp[i][j] = true
				} else {
					predp[i][j] = predp[i+1][j-1]
				}
			}
		}
	}

	dp := make([]int, n)
	for k, _ := range dp {
		dp[k] = k
	}
	//for k,_:=range dp{
	//	dp[k]=make([]int,n)
	//}

	for i := 1; i < n; i++ {
		if predp[0][i] {
			dp[i] = 0
			continue
		}

		for j := 0; j < i; j++ {
			if predp[j+1][i] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	//for i:=n-1;i>=0;i--{
	//	for j:=i;j<n;j++{
	//		if predp[i][j]{
	//			dp[i][j]=0
	//		}else{
	//			dp[i][j]=j-i+1
	//			for k:=i;k<j;k++{
	//				dp[i][j]=min(dp[i][j],dp[i][k]+dp[k+1][j]+1)
	//			}
	//		}
	//	}
	//}
	return dp[n-1]
}
