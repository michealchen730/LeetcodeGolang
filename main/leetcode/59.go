package main

func generateMatrix(n int) [][]int {
	if n==0 {
		return [][]int{}
	}

	res:=make([][]int,n)
	for i:=0; i<n; i++ {
		res[i]=make([]int,n)
	}

	//x表示从第x行第x列开始
	//t表示第t-1行第t-1结束
	x,t,last:=0,n,n*n
	temp:=1

	for temp<=last{
		for s:=x;s<t;s++{
			res[x][s]=temp
			temp++
		}
		for s:=x+1;s<t;s++{
			res[s][t-1]=temp
			temp++
		}
		for s:=t-2;s>=x;s--{
			res[t-1][s]=temp
			temp++
		}
		for s:=t-2;s>x;s--  {
			res[s][x]=temp
			temp++
		}
		x++
		t--
	}
	return res
}
