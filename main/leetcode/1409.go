package main

func processQueries(queries []int, m int) []int {
	var res []int
	p := make([]int, m)
	for k, _ := range p {
		p[k] = k + 1
	}
	for _, v := range queries {
		i := 0
		for p[i] != v {
			i++
		}
		res = append(res, i)
		p = append(p[:i], p[i+1:]...)
		p = append([]int{v}, p...)
	}
	return res
}
