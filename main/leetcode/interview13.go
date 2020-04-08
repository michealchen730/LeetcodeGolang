package main

func movingCount(m int, n int, k int) int {
	res := 1
	mat := make([][]int, m)
	for r, _ := range mat {
		mat[r] = make([]int, n)
	}
	stack := [][]int{{0, 0}}
	for len(stack) != 0 {
		temp := stack[0]
		if temp[0]+1 < m && mat[temp[0]+1][temp[1]] != 1 {
			if getCount(temp[0]+1, temp[1]) <= k {
				stack = append(stack, []int{temp[0] + 1, temp[1]})
				res++
			}
			mat[temp[0]+1][temp[1]] = 1
		}
		if temp[1]+1 < n && mat[temp[0]][temp[1]+1] != 1 {
			if getCount(temp[0], temp[1]+1) <= k {
				stack = append(stack, []int{temp[0], temp[1] + 1})
				res++
			}
			mat[temp[0]][temp[1]+1] = 1
		}
		stack = stack[1:]
	}
	return res
}
func getCount(ti, tj int) int {
	count := 0
	for ti != 0 || tj != 0 {
		count += ti%10 + tj%10
		ti /= 10
		tj /= 10
	}
	return count
}
