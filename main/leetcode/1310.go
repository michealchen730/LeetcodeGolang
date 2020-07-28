package main

func xorQueries(arr []int, queries [][]int) []int {
	sum := make([]int, len(arr)+1)
	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] ^ arr[i-1]
	}
	res := make([]int, len(queries))
	for k, v := range queries {
		res[k] = sum[v[0]] ^ sum[v[1]+1]
	}
	return res
}
