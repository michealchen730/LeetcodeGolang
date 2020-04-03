package main

func sumZero(n int) []int {
	var res []int
	if n%2 != 0 {
		res = append(res, 0)
	}
	i := 1
	for len(res) != n {
		res = append(res, []int{i, -i}...)
		i++
	}
	return res
}
