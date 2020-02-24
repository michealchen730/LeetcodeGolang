package main

func getMoneyAmount(n int) int {
	res := make([][]int, n+1)
	for k, _ := range res {
		res[k] = make([]int, n+1)
	}
	for i := n - 1; i > 0; i-- {
		for j := i + 1; j < len(res); j++ {
			temp := res[i][j-1] + j
			for k := i; k < j; k++ {
				t := max(res[i][k-1], res[k+1][j]) + k
				if t < temp {
					temp = t
				}
			}
			res[i][j] = temp
		}
	}
	return res[1][n]
}

func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
