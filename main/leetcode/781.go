package main

func numRabbits(answers []int) int {
	res := 0
	mp := make(map[int]int)
	for _, v := range answers {
		//第一个为v的
		if mp[v] == 0 {
			res += v + 1
		}
		mp[v]++
		if mp[v] == v+1 {
			mp[v] = 0
		}
	}
	return res
}
