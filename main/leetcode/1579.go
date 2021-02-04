package main

func maxNumEdgesToRemove(n int, edges [][]int) int {
	root := make([]int, n+1)
	rank := make([]int, n+1)
	for i := 1; i <= n; i++ {
		root[i] = i
		rank[i] = 1
	}
	root2, rank2 := make([]int, n+1), make([]int, n+1)
	copy(root2, root)
	copy(rank2, rank)
	n1, n2 := n, n
	res := 0
	for _, v := range edges {
		if v[0] == 3 {
			if Cha(root, v[1]) == Cha(root, v[2]) {
				res += 1
			} else {
				Bing(root, rank, v[1], v[2])
				Bing(root2, rank2, v[1], v[2])
				n1 -= 1
				n2 -= 1
			}
		}
	}
	for _, v := range edges {
		if v[0] == 1 {
			if Cha(root, v[1]) == Cha(root, v[2]) {
				res += 1
			} else {
				Bing(root, rank, v[1], v[2])
				n1 -= 1
			}
		}
		if v[0] == 2 {
			if Cha(root2, v[1]) == Cha(root2, v[2]) {
				res += 1
			} else {
				Bing(root2, rank2, v[1], v[2])
				n2 -= 1
			}
		}
	}
	if n1 != 1 || n2 != 1 {
		return -1
	}
	return res
}
