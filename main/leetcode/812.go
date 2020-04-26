package main

func largestTriangleArea(points [][]int) float64 {
	res := 0
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				temp := points[i][0]*points[j][1] + points[j][0]*points[k][1] + points[k][0]*points[i][1] - points[j][1]*points[k][0] - points[i][1]*points[j][0] - points[i][0]*points[k][1]
				if abs(temp) > res {
					res = abs(temp)
				}
			}
		}
	}
	return float64(res) / 2
}
