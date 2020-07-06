package main

func stoneGameIII(stoneValue []int) string {
	dp := make([]int, len(stoneValue))
	sum := stoneValue[len(dp)-1]
	dp[len(dp)-1] = sum
	for i := len(dp) - 2; i >= 0; i-- {
		sum += stoneValue[i]
		dp[i] = sum - dp[i+1]
		v := 0
		for j := 1; j < 3 && j+i < len(dp); j++ {
			v += stoneValue[i+j]
			if i+j+1 < len(dp) {
				dp[i] = max(dp[i], sum-dp[i+j+1])
			} else {
				dp[i] = max(dp[i], sum)
			}
		}
	}
	if dp[0] > sum-dp[0] {
		return "Alice"
	} else if dp[0] < sum-dp[0] {
		return "Bob"
	} else {
		return "Tie"
	}
}
