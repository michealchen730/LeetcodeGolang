package main

func pancakeSort(A []int) []int {
	var res []int
	for k := len(A); k >= 1; k-- {
		for i, v := range A {
			if v == k && i != k-1 {
				if i != 0 {
					reverseArr(A[:i+1])
					res = append(res, i+1)
				}
				reverseArr(A[:k])
				res = append(res, k)
			}
		}
	}
	return res
}
