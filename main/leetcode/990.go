package main

func equationsPossible(equations []string) bool {
	arr := make([]int, 26)
	for k, _ := range arr {
		arr[k] = k
	}
	var queue []int
	for k, v := range equations {
		if v[1] == '=' {
			UnionFriends990(arr, int(v[0]-'a'), int(v[3]-'a'))
		} else {
			queue = append(queue, k)
		}
	}
	for _, v := range queue {
		if findUnion990(arr, int(equations[v][0]-'a')) == findUnion990(arr, int(equations[v][3]-'a')) {
			return false
		}
	}
	return true
}

func UnionFriends990(stu []int, i, j int) int {
	a, b := findUnion990(stu, i), findUnion990(stu, j)
	if a != b {
		stu[b] = a
		return 1
	}
	return 0
}

func findUnion990(stu []int, i int) int {
	for stu[i] != i {
		i = stu[i]
	}
	return i
}
