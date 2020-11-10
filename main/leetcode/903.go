package main

//dp[i][j]表示前i位符合其中最后一位为j的个数
func numPermsDISequence(S string) int {
	N := len(S)

	dp := make([][]int, N+1)
	for k, _ := range dp {
		dp[k] = make([]int, N+1)
	}
	//当序列长度为1，那么拼出来的只能为“0”
	for i := 0; i <= N; i++ {
		dp[0][i] = 1
	}

	//S的长度为N，那么就有N+1个数字
	for i := 1; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if S[i-1] == 'D' {
				for k := j; k < i; k++ {
					dp[i][j] += dp[i-1][k]
					dp[i][j] %= 1000000007
				}
			} else {
				for k := 0; k < j; k++ {
					dp[i][j] += dp[i-1][k]
					dp[i][j] %= 1000000007
				}
			}
		}
	}

	res := 0
	for _, v := range dp[N] {
		res += v
		res %= 1000000007
	}
	return res
}
