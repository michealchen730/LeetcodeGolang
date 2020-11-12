package main

import "math"

func findRotateSteps(ring string, key string) int {
	dp := make([][]int, len(key)+1)
	for k, _ := range dp {
		dp[k] = make([]int, len(ring)+1)
		for i, _ := range dp[k] {
			dp[k][i] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	for i := 1; i <= len(key); i++ {
		for j := 1; j <= len(ring); j++ {
			//如果相等
			if key[i-1] == ring[j-1] {
				if i == 1 {
					dp[i][j] = min(j-1, 1+len(ring)-j) + 1
					continue
				}
				for m, v := range dp[i-1] {
					if v < math.MaxInt32 {
						length := len(ring)
						if m < j {
							length = min(j-m, m+len(ring)-j)
						} else {
							length = min(m-j, j+len(ring)-m)
						}
						dp[i][j] = min(dp[i][j], dp[i-1][m]+length+1)
					}
				}
			}
		}
	}

	res := math.MaxInt32
	for _, v := range dp[len(key)] {
		res = min(res, v)
	}
	return res
}

//可以使用map做一轮优化，不需要再遍历二维数组的所有位置
//用了map反而更慢了（数据结构本身的问题，可以考虑用另一轮二维数组去代替）
func findRotateSteps2(ring string, key string) int {
	mp := make(map[byte][]int)
	for k, _ := range ring {
		mp[ring[k]] = append(mp[ring[k]], k+1)
	}

	dp := make([][]int, len(key)+1)
	for k, _ := range dp {
		dp[k] = make([]int, len(ring)+1)
		for i, _ := range dp[k] {
			dp[k][i] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	for i := 1; i <= len(key); i++ {
		for _, v := range mp[key[i-1]] {
			if i == 1 {
				dp[i][v] = min(v-1, 1+len(ring)-v) + 1
				continue
			}
			for _, x := range mp[key[i-2]] {
				length := len(ring)
				if x < v {
					length = min(v-x, x+len(ring)-v)
				} else {
					length = min(x-v, v+len(ring)-x)
				}
				dp[i][v] = min(dp[i][v], dp[i-1][x]+length+1)
			}
		}
	}

	res := math.MaxInt32
	for _, v := range mp[key[len(key)-1]] {
		res = min(res, dp[len(key)][v])
	}
	return res

}
