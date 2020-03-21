package main

func groupThePeople(groupSizes []int) [][]int {
	mp := make(map[int][]int)
	var res [][]int
	for k, v := range groupSizes {
		mp[v] = append(mp[v], k)
		if len(mp[v])%v == 0 {
			res = append(res, mp[v][len(mp[v])-v:])
		}
	}
	return res
}
