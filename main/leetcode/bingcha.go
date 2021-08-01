package main

func Bing(stu []int, rank []int, i, j int) {
	a, b := Cha(stu, i), Cha(stu, j)
	if a == b {
		return
	}
	if rank[a] < rank[b] {
		a, b = b, a
	}
	rank[a] += rank[b]
	stu[b] = a
}

func Cha(stu []int, i int) int {
	for stu[i] != i {
		i = stu[i]
	}
	return i
}
