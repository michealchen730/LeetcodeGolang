package main

func countArrangement(N int) int {
	cnt := 0
	used := make([]bool, N+1)
	var dfs func(n int)

	dfs = func(n int) {
		if n == N+1 {
			cnt++
			return
		}
		flag := true
		for i := 1; i <= N; i++ {
			if !used[i] && (i%n == 0 || n%i == 0) {
				flag = true
				used[i] = true
				dfs(n + 1)
				used[i] = false
			}
		}
		if flag {
			return
		}
	}
	dfs(1)
	return cnt
}
