package main

func restoreArray(adjacentPairs [][]int) []int {
	var res []int
	if len(adjacentPairs) == 0 {
		return res
	}
	res = make([]int, len(adjacentPairs)+1)
	mp := make(map[int][]int)
	for _, v := range adjacentPairs {
		mp[v[0]] = append(mp[v[0]], v[1])
		mp[v[1]] = append(mp[v[1]], v[0])
	}
	for k, v := range mp {
		if len(v) == 1 {
			res[0] = k
			break
		}
	}
	tmp := 0
	for tmp < len(adjacentPairs) {
		t := res[tmp]
		cand := mp[t]
		tmp++
		if len(cand) == 1 {
			res[tmp] = cand[0]
		} else {
			if res[tmp-2] == cand[0] {
				res[tmp] = cand[1]
			} else {
				res[tmp] = cand[0]
			}
		}
	}
	return res
}
