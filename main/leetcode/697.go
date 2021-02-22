package main

func findShortestSubArray(nums []int) int {
	mx := 0
	res := len(nums)
	cnt := make(map[int]int)
	start := make(map[int]int)
	for k, v := range nums {
		cnt[v]++
		if cnt[v] > mx {
			res = k - start[v] + 1
			mx = cnt[v]
		}
		if cnt[v] == mx {
			res = min(res, k-start[v]+1)
		}
		if cnt[v] == 1 {
			start[v] = k
		}
	}
	return res
}
