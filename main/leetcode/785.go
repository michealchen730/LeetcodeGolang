package main

func isBipartite(graph [][]int) bool {
	//测试的用例中给出了某个顶点没有任何边的选项，可以理解为有两个图，这两个图并不相连
	//可以将此用例扩开，如果有三个图，且互不相连，那么简单的bfs一定不可行
	bfs, n := 0, len(graph) //bfs这个变量表示我们遍历过的顶点数目
	mp1, mp2 := make(map[int]bool), make(map[int]bool)
	used := make([]bool, len(graph))
	var queue []int
	for bfs != n {
		for k, v := range used {
			if !v {
				used[k] = true
				queue = []int{k}
				mp1[k] = true
				for len(queue) != 0 {
					t := queue[0]
					bfs++
					if mp1[t] {
						for _, v := range graph[t] {
							if mp1[v] {
								return false
							}
							mp2[v] = true
						}
					} else {
						for _, v := range graph[t] {
							if mp2[v] {
								return false
							}
							mp1[v] = true
						}
					}
					for _, v := range graph[t] {
						if !used[v] {
							used[v] = true
							queue = append(queue, v)
						}
					}
					queue = queue[1:]
				}
			}
		}
	}
	return true
}
