package main

func findCircleNum(M [][]int) int {
	res := len(M)
	students := make([]int, len(M))
	for k, _ := range students {
		students[k] = k
	}
	for i := 0; i < len(M); i++ {
		for j := i; j < len(M); j++ {
			if M[i][j] == 1 {
				res -= UnionFriends(students, i, j)
			}
		}
	}
	return res
}

func UnionFriends(stu []int, i, j int) int {
	a, b := findUnion(stu, i), findUnion(stu, j)
	if a != b {
		stu[b] = a
		return 1
	}
	return 0
}

func findUnion(stu []int, i int) int {
	for stu[i] != i {
		i = stu[i]
	}
	return i
}
