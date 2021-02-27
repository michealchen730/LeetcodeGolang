package main

func numsSameConsecDiff(n int, k int) []int {
	var res []int
	var path []int
	var dfs func()
	dfs = func() {
		if len(path) == n {
			res = append(res, arrToInt(path))
			return
		}
		if len(path) == 0 {
			for i := 1; i <= 9; i++ {
				path = append(path, i)
				dfs()
				path = path[:len(path)-1]
			}
		} else {
			t := path[len(path)-1]
			if t-k >= 0 && t-k <= 9 {
				path = append(path, t-k)
				dfs()
				path = path[:len(path)-1]
			}
			if t+k >= 0 && t+k <= 9 && k != 0 {
				path = append(path, t+k)
				dfs()
				path = path[:len(path)-1]
			}
		}
	}
	dfs()
	return res
}
