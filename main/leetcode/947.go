package main

func removeStones(stones [][]int) int {
	arr := make([]int, len(stones))
	for i, _ := range arr {
		arr[i] = i
	}
	rowMp := make(map[int]int)
	colMp := make(map[int]int)
	for k, v := range stones {
		_, ok1 := rowMp[v[0]]
		_, ok2 := colMp[v[1]]
		if ok1 {
			UnionStones947(arr, rowMp[v[0]], k)
		} else {
			rowMp[v[0]] = k
		}
		if ok2 {
			UnionStones947(arr, colMp[v[1]], k)
		} else {
			rowMp[v[1]] = k
		}
	}
	res := 0
	for k, v := range arr {
		if k != v {
			res++
		}
	}
	return res
}

func UnionStones947(arr []int, i, j int) int {
	a, b := findUnion947(arr, i), findUnion947(arr, j)
	if a != b {
		arr[b] = a
		return 1
	}
	return 0
}

func findUnion947(arr []int, i int) int {
	for arr[i] != i {
		i = arr[i]
	}
	return i
}
