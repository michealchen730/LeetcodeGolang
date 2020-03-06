package main

func findContinuousSequence(target int) [][]int {
	if target == 3 {
		return [][]int{{1, 2}}
	}
	var res [][]int
	for j := target / 2; j > 1; j-- {
		if j%2 != 0 && target%j == 0 && target/j > j/2 {
			arr := make([]int, j)
			for i := 0; i < len(arr); i++ {
				arr[i] = target/j - j/2 + i
			}
			res = append(res, arr)
		}
		if j%2 == 0 && target%j == j/2 && target/j > j/2-1 {
			arr := make([]int, j)
			for i := 0; i < len(arr); i++ {
				arr[i] = target/j - j/2 + 1 + i
			}
			res = append(res, arr)
		}
	}
	return res
}
