package main

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for k := 1; k < len(points); k++ {
		d1, d2 := abs(points[k][0]-points[k-1][0]), abs(points[k][1]-points[k-1][1])
		step1 := min(d1, d2)
		step2 := abs(d1 - d2)
		res += step1 + step2
	}
	return res
}
