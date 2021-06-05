package main

import "fmt"

func dpali(arr []int, l int) int {
	sum := make([]int, l+1)
	for i := 1; i <= l; i++ {
		sum[i] = sum[i-1] + arr[i-1]
	}

	dp := make([][]int, l)
	for k, _ := range dp {
		dp[k] = make([]int, l)
		dp[k][k] = arr[k]
	}

	for j := l - 1; j >= 0; j-- {
		for i := j; i < l; i++ {
			if i == j {
				continue
			}
			if j == 2 && i == 5 {
				fmt.Println("catch!!")
			}
			for k := j + 1; k <= i; k++ {
				if sum[i+1]-sum[k] > sum[k]-sum[j] {
					temp := sum[k] - sum[j]
					if j != k-1 {
						temp += dp[j][k-1]
					}
					dp[j][i] = max(dp[j][i], temp)
				} else {
					temp := sum[i+1] - sum[k]
					if k != i {
						temp += dp[k][i]
					}
					dp[j][i] = max(dp[j][i], temp)
				}
			}
		}
	}
	return dp[0][l-1]
}
