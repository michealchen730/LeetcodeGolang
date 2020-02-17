package main

func gardenNoAdj(N int, paths [][]int) []int {
	if N==0 {
		return nil
	}
	//先遍历整个paths
	//还是需要一个长度为N+1的数组
	//存在没有出现在path中的点，即单独的花园，数组上这些点全部赋值为1
	//还是用一个map去存放邻接表，不要使用数组，万一大部分花园都是单独存在的，会浪费空间
	//遍历数组，当数组元素为0，去map中找这个元素的邻接表，看下那几个元素被用了，判断元素被用的点，我想也用一个used数组来表示，长度为4
	used:=make([]int,N+1)
	vertex:=make(map[int][]int)
	for _,v:=range paths{
		vertex[v[0]]=append(vertex[v[0]],v[1])
		vertex[v[1]]=append(vertex[v[1]],v[0])
		used[v[0]],used[v[1]]=-1,-1
	}
	flowers:=[5]int{0,0,0,0,0}
	for i:=1; i<=N; i++{
		if used[i]!=-1 {
			used[i]=1
			continue
		}
		for j:=0;j<len(vertex[i]);j++{
			if used[vertex[i][j]]!=-1 {
				flowers[used[vertex[i][j]]]=1
			}
		}
		for j:=1;j<len(flowers);j++{
			if flowers[j]==0 {
				used[i]=j
				break
			}
		}
		for j:=1;j<=4;j++{
			flowers[j]=0
		}
	}
	res:=used[1:]
	return res
}
