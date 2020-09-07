package main

func mostVisited(n int, rounds []int) []int {
	start, end := rounds[0], rounds[len(rounds)-1]
	var res []int
	if start <= end {
		for k := start; k <= end; k++ {
			res = append(res, k)
		}
	} else {
		for k := 1; k <= end; k++ {
			res = append(res, k)
		}
		for k := start; k <= n; k++ {
			res = append(res, k)
		}
	}
	return res
}
