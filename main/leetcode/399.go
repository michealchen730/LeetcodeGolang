package main

//转化为图的形式
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	res := make([]float64, len(queries))
	mp := make(map[string]int)
	F := 1
	for _, v := range equations {
		for _, word := range v {
			if mp[word] == 0 {
				mp[word] = F
				F++
			}
		}
	}

	mat := make([][]float64, F+1)
	for k, _ := range mat {
		mat[k] = make([]float64, F+1)
	}

	for k, v := range equations {
		a1, a2 := mp[v[0]], mp[v[1]]
		mat[a1][a2] = values[k]
		mat[a2][a1] = 1 / values[k]
	}
	var findRes func([]string) float64
	findRes = func(strs []string) float64 {
		if mp[strs[0]] == 0 || mp[strs[1]] == 0 {
			return -1
		}
		if strs[0] == strs[1] {
			return 1
		}
		used := make([]bool, F+1)
		query := []int{mp[strs[0]]}
		vals := []float64{1}
		used[mp[strs[0]]] = true
		target := mp[strs[1]]
		for len(query) != 0 {
			l := len(query)
			for i := 0; i < l; i++ {
				idx := query[i]
				mult := vals[i]
				for j := 1; j <= F; j++ {
					if mat[idx][j] != 0 && !used[j] {
						if j == target {
							return mult * mat[idx][j]
						} else {
							query = append(query, j)
							vals = append(vals, mult*mat[idx][j])
							used[j] = true
						}
					}
				}
			}
			query = query[l:]
			vals = vals[l:]
		}
		return -1
	}

	for k, v := range queries {
		res[k] = findRes(v)
	}

	return res
}
