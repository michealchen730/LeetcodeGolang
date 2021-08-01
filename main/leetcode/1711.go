package main

import "fmt"

func countPairs(deliciousness []int) int {
	MOD := 1000000007
	powers := make([]int, 31)
	for k, _ := range powers {
		powers[k] = 1 << k
	}
	mp := make(map[int]int)
	for _, v := range deliciousness {
		mp[v]++
	}
	res := 0
	for k, v := range mp {
		for _, p := range powers {
			if p >= k && mp[p-k] != 0 {
				if p-k == k {
					res += v * (v - 1) / 2
				} else {
					res += v * mp[p-k]
				}
				res %= MOD
			}
		}
		delete(mp, k)
	}
	return res
}

func main() {
	fmt.Println(countPairs([]int{0, 1}))
}
