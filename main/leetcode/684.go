package main

func findRedundantConnection(edges [][]int) []int {
	arr := make([]int, len(edges))
	for k, _ := range arr {
		arr[k] = k
	}
	for _, v := range edges {
		if findUnion(arr, v[0]) != findUnion(arr, v[1]) {
			UnionVertex(arr, v[0], v[1])
		} else {
			return []int{v[0], v[1]}
		}
	}
	return []int{}
}

func UnionVertex(ver []int, i, j int) int {
	a, b := findUnion(ver, i), findUnion(ver, j)
	if a != b {
		ver[b] = a
		return 1
	}
	return 0
}
