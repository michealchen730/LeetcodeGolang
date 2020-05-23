package main

func fairCandySwap(A []int, B []int) []int {
	res := []int{0, 0}
	mp := make(map[int]int)
	sum1, sum2 := 0, 0
	for _, v := range A {
		sum1 += v
	}
	for _, v := range B {
		sum2 += v
		mp[v]++
	}
	d := sum1 - sum2
	for _, v := range A {
		if (v*2-d)%2 == 0 && mp[(v*2-d)/2] != 0 {
			res[0] = v
			res[1] = (v*2 - d) / 2
			break
		}
	}
	return res
}
