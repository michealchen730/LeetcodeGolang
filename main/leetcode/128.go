package main

func longestConsecutive(nums []int) int {
	mp := make(map[int]int)
	res := 0
	for _, v := range nums {
		if mp[v] == 0 {
			mp[v] = mp[v-1] + mp[v+1] + 1
			if mp[v] > res {
				res = mp[v]
			}
			mp[v-mp[v-1]] = mp[v]
			mp[v+mp[v+1]] = mp[v]
		}
	}
	return res
}
