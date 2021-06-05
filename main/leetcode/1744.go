package main

func canEat(candiesCount []int, queries [][]int) []bool {
	arr := make([]int, len(candiesCount)+1)
	res := make([]bool, len(queries))
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] + candiesCount[i-1]
	}
	for i, v := range queries {
		mn, mx := v[1]+1, v[2]*(v[1]+1)
		res[i] = canOverlap(mn, mx, arr[v[0]]+1, arr[v[0]+1])
	}
	return res
}

func canOverlap(l1, r1, l2, r2 int) bool {
	if r1 < l2 || l1 > r2 {
		return false
	}
	return true
}
