package main

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	m := 1000000007
	nums := len(group)
	profits := 0
	for _, v := range profit {
		profits += v
	}
	dp := make([][]int, profits+1)
	for k, _ := range dp {
		dp[k] = make([]int, n+1)
	}
	dp[0][0] = 1
	for idx := 1; idx <= nums; idx++ {
		for prf := profits; prf >= 0; prf-- {
			for prs := n; prs >= 0; prs-- {
				if prf >= profit[idx-1] && prs >= group[idx-1] {
					dp[prf][prs] = (dp[prf][prs] + dp[prf-profit[idx-1]][prs-group[idx-1]]) % m
				}
			}
		}
	}
	res := 0
	for p := minProfit; p <= profits; p++ {
		for j := 0; j <= n; j++ {
			res += dp[p][j]
			res %= m
		}
	}
	return res
}
