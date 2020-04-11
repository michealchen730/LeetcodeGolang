package main

func superEggDrop(K int, N int) int {
	if K == 1 {
		return N
	}
	mat := make([][]int, K+1)
	for k, _ := range mat {
		mat[k] = make([]int, N+1)
	}
	for k, _ := range mat[1] {
		mat[1][k] = k
	}
	for i := 2; i <= K; i++ {
		for j := 1; j <= N; j++ {
			low, high := 1, j
			for low+1 < high {
				mid := (low + high) / 2
				if mat[i-1][mid-1] < mat[i][j-mid] {
					low = mid
				} else if mat[i-1][mid-1] > mat[i][j-mid] {
					high = mid
				} else {
					low, high = mid, mid
				}
			}
			mat[i][j] = 1 + min(max(mat[i-1][low-1], mat[i][j-low]), max(mat[i-1][high-1], mat[i][j-high]))
		}
	}
	return mat[K][N]
}
