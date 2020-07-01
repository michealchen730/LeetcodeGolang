package main

func average(salary []int) float64 {
	mn, mx := 10000000, -1
	res := 0
	for _, v := range salary {
		res += v
		mn = min(mn, v)
		mx = max(mx, v)
	}
	return float64(res-mx-mn) / float64(len(salary)-2)
}
