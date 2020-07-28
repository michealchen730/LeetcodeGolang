package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 || len(s2) == 0 {
		return s1 == s3 || s2 == s3
	}
	dp := make([][]bool, len(s3))
	for k, _ := range dp {
		dp[k] = make([]bool, len(s1)+1)
	}
	if s1[0] == s3[0] {
		dp[0][1] = true
	}
	if s2[0] == s3[0] {
		dp[0][0] = true
	}
	for i := 1; i < len(s3); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if dp[i-1][j] == true {
				if j < len(s1) && s1[j] == s3[i] {
					dp[i][j+1] = true
				}
				if i-j < len(s2) && s2[i-j] == s3[i] {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[len(s3)-1][len(s1)+1]
}
