package main

//标签为深度优先搜索
//在我看来更像是拓扑排序
//func loudAndRich(richer [][]int, quiet []int) []int {
//	//将关系转化为图
//	//需要一个arr来记录出入度
//	n:=len(quiet)
//	arr := make([]int,n)
//	mat := make([][]int,n)
//	for k,_:=range mat{
//		mat[k] = make([]int,n)
//	}
//	for _,v:=range richer{
//		mat[v[0]][v[1]]=1
//		arr[v[1]]++
//	}
//	//创建一个二维数组mat2，mat2[i]记录钱不少于i的名单，初始化是mat2[i] = [i]
//	mat2:=make([]map[int]int,n)
//	for k,_:=range mat2{
//		mat2[k] = map[int]int{k:1}
//	}
//	//创建结果集数组
//	res:=make([]int,n)
//	//逐项搜索
//	solved:=0
//	for solved<n{
//		for k,v:=range arr{
//			//如果入度为0
//			if v == 0{
//				arr[k]--
//				solved++
//				q:=quiet[k]
//				res[k]=k
//				//遍历比k富有的人的列表，找出谁最安静并保存
//				for v2,_:=range mat2[k]{
//					if quiet[v2]<q{
//						q=quiet[v2]
//						res[k]=v2
//					}
//				}
//				//处理与k相关的边
//				for k3,v3:=range mat[k]{
//					if v3 == 1{
//						//将比k富有的人的列表加入到没有k富有的人中
//						for key,_:=range mat2[k]{
//							mat2[k3][key]++
//						}
//						//穷人的入度-1
//						arr[k3]--
//					}
//				}
//			}
//		}
//
//	}
//	return res
//}
//树形DFS
//func loudAndRich(richer [][]int, quiet []int) []int {
//	//将关系转化为图
//	//需要一个arr来记录出入度
//	n:=len(quiet)
//	mat := make([][]int,n)
//	for k,_:=range mat{
//		mat[k] = make([]int,n)
//	}
//	for _,v:=range richer{
//		mat[v[1]][v[0]]=1
//	}
//	//创建结果集数组
//	res:=make([]int,n)
//	for k,_:=range res{
//		res[k]=-1
//	}
//
//	var dfs func(idx int) int
//	dfs = func(idx int) int {
//		if res[idx]==-1{
//			res[idx]=idx
//			for k,v:=range mat[idx]{
//				if v==1&&quiet[dfs(k)]<quiet[res[idx]]{
//					res[idx] = dfs(k)
//				}
//			}
//		}
//		return res[idx]
//	}
//	for k,_:=range res{
//		dfs(k)
//	}
//	return res
//}
//降低空间复杂度的写法，但时间复杂度更高了
func loudAndRich(richer [][]int, quiet []int) []int {
	//将关系转化为图
	//需要一个arr来记录出入度
	n := len(quiet)
	mat := make([][]int, n)
	for k, _ := range mat {
		mat[k] = []int{}
	}
	for _, v := range richer {
		mat[v[1]] = append(mat[v[1]], v[0])
	}
	//创建结果集数组
	res := make([]int, n)
	for k, _ := range res {
		res[k] = -1
	}
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if res[idx] == -1 {
			res[idx] = idx
			for _, v := range mat[idx] {
				if quiet[dfs(v)] < quiet[res[idx]] {
					res[idx] = dfs(v)
				}
			}
		}
		return res[idx]
	}
	for k, _ := range res {
		dfs(k)
	}
	return res
}
