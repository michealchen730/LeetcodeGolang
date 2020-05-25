package main

func escapeGhosts(ghosts [][]int, target []int) bool {
	mn := abs(target[0]) + abs(target[1])
	for _, v := range ghosts {
		if abs(v[0]-target[0])+abs(v[1]-target[1]) <= mn {
			return false
		}
	}
	return true
}
