package main

func subarraysDivByK(A []int, K int) int {
	rec := map[int]int{0: 1}
	sum, res := 0, 0
	for _, v := range A {
		sum += v
		mod := (sum%K + K) % K //-1%3=-1 -5%3=-2 (sum%K+K)%K主要是应对负数出现的情况
		res += rec[mod]
		rec[mod]++
	}
	return res
}
