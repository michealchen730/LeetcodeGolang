package main

func oddCells(n int, m int, indices [][]int) int {
	arr1, arr2 := make([]int, n), make([]int, m)
	//统计有效行操作和有效列操作
	for _, v := range indices {
		arr1[v[0]]++
		arr2[v[1]]++
		if arr1[v[0]] >= 2 {
			arr1[v[0]] -= 2
		}
		if arr2[v[1]] >= 2 {
			arr2[v[1]] -= 2
		}
	}
	rows := 0 //代表所有行操作后，有多少行是偶数，这些行会在列操作后最终变成奇数
	for _, v := range arr1 {
		if v == 0 {
			rows++
		}
	}
	sum := 0
	for _, v := range arr2 {
		if v == 1 {
			//有列操作，sum+=偶数行
			sum += rows
		} else {
			//无列操作，sum+=奇数行
			sum += n - rows
		}
	}
	return sum
}
