package main

import "fmt"

func eventualSafeNodes(graph [][]int) []int {
	arr := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if arr[i] == 0 {
			var path []int
			dfs802(i, graph, arr, &path)
		}
	}
	var res []int
	for k, v := range arr {
		if v != 1 {
			res = append(res, k)
		}
	}
	return res
}

func dfs802(tmp int, graph [][]int, arr []int, path *[]int) {
	fmt.Print("dfs")
	fmt.Println(tmp)
	*path = append(*path, tmp)
	if len(graph[tmp]) == 0 || arr[tmp] == -1 {
		//这里可以再加一点东西，但是需要改变arr数组的格式
		*path = (*path)[:len(*path)-1]
		return
	}
	if arr[tmp] == 1 || arr[tmp] == 2 { //如果通向已经证实过的环
		*path = (*path)[:len(*path)-1]
		for i := 0; i < len(*path); i++ {
			arr[(*path)[i]] = 1
		}
		return
	} else {
		//此时需要验证是否有环
		arr[tmp] = 2
		for i := 0; i < len(graph[tmp]); i++ {
			dfs802(graph[tmp][i], graph, arr, path)
			if arr[tmp] == 1 {
				break
			}
		}
		if arr[tmp] != 1 {
			arr[tmp] = -1
		}
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	fmt.Println(eventualSafeNodes([][]int{[]int{1, 2}, []int{2, 3}, []int{5}, []int{0}, []int{5}, []int{6}, []int{}}))
}
