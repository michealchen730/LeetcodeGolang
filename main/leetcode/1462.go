package main

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	mp := make(map[int][]int)
	for _, v := range prerequisites {
		mp[v[1]] = append(mp[v[1]], v[0])
	}
	res := make([]bool, len(queries))
	for k, v := range queries {
		used := make([]int, n)
		queue := []int{v[1]}
		for len(queue) != 0 {
			t := queue[0]
			if t == v[0] {
				res[k] = true
				break
			}
			queue = queue[1:]
			for _, v := range mp[t] {
				if used[v] == 0 {
					queue = append(queue, v)
					used[v] = 1
				}
			}
		}
	}
	return res
}
