package main

func maxScoreSightseeingPair(A []int) int {
	mx := A[0] + 0
	res := A[0] + A[1] - 1
	for k, v := range A {
		if k > 0 {
			res = max(res, mx+v-k)
			mx = max(mx, k+v)
		}
	}
	return res
}
