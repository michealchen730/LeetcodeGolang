package main

func pancakeSort(A []int) []int {
	var res []int
	for k := len(A); k >= 1; k-- {
		for i, v := range A {
			if v == k && i != k-1 {
				if i != 0 {
					reverse(A[:i+1])
					res = append(res, i+1)
				}
				reverse(A[:k])
				res = append(res, k)
			}
		}
	}
	return res
}
