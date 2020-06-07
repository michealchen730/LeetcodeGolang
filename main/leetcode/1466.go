package main

func minReorder(n int, connections [][]int) int {
	res := 0
	mp1 := make(map[int][]int)
	mp2 := make(map[int][]int)
	for _, v := range connections {
		mp1[v[0]] = append(mp1[v[0]], v[1])
		mp2[v[1]] = append(mp2[v[1]], v[0])
	}
	queue := []int{0}
	used := make([]int, n)
	used[0] = 1
	for len(queue) != 0 {
		t := queue[0]
		queue = queue[1:]
		for _, v := range mp2[t] {
			if used[v] == 0 {
				queue = append(queue, v)
				used[v] = 1
			}
		}
		for _, v := range mp1[t] {
			if used[v] == 0 {
				queue = append(queue, v)
				used[v] = 1
				res++
			}
		}
	}
	return res
}
