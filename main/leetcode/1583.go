package main

//func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
//	pairsMap:=make(map[int]int)
//	for _,v:=range pairs{
//		pairsMap[v[0]]=v[1]
//		pairsMap[v[1]]=v[0]
//	}
//	matrix:=make([][][]int, n)
//	for i,v:=range preferences{
//		for k,val:=range v {
//			matrix[i] = append(matrix[i], []int{val, k})
//		}
//		matrix[i] = append(matrix[i], []int{i,i})
//		sort.Slice(matrix[i], func(a, b int) bool {
//			return matrix[i][a][0]<matrix[i][b][0]
//		})
//	}
//	res:=0
//	for i:=0;i<n;i++{
//		partner:=pairsMap[i]
//		for _,rival:=range preferences[i]{
//			if rival==partner{
//				break
//			}else{
//				rivalPartner:=pairsMap[rival]
//				if matrix[rival][i][1]<matrix[rival][rivalPartner][1]{
//					res++
//					break
//				}
//			}
//		}
//	}
//	return res
//}

//优化了，省了排序
func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	pairsMap := make(map[int]int)
	for _, v := range pairs {
		pairsMap[v[0]] = v[1]
		pairsMap[v[1]] = v[0]
	}
	matrix := make([][]int, n)
	for i, v := range preferences {
		matrix[i] = make([]int, n)
		for k, val := range v {
			matrix[i][val] = k
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		partner := pairsMap[i]
		for _, rival := range preferences[i] {
			if rival == partner {
				break
			} else {
				rivalPartner := pairsMap[rival]
				if matrix[rival][i] < matrix[rival][rivalPartner] {
					res++
					break
				}
			}
		}
	}
	return res
}
