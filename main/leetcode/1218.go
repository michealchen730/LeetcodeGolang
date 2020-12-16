package main

func longestSubsequence(arr []int, difference int) int {
	mp := make(map[int]int)
	res := 0

	for _, v := range arr {
		mp[v] = mp[v-difference] + 1
		res = max(res, mp[v])
	}

	return res
}
