package main

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix)==0 {
		return res
	}
	length,width,i,j:=len(matrix)-1,len(matrix[0])-1,0,0
	for i<=length&&j<=width {
		getBorder(matrix,&res,i,j,length,width)
		i++
		j++
		length--
		width--
	}
	return res
}

func getBorder(matrix [][]int,res *[]int,a int,b int,c int,d int)  {
	//针对nxn矩阵的最内层
	if a==c&&b==d {
		*res=append(*res, matrix[a][b])
		return
	}
	//针对mxn矩阵的最内层：m>n情况下，且最内层是单行
	if a==c&&b!=d{
		for i:=b; i<=d; i++{
			*res=append(*res, matrix[a][i])
		}
		return
	}
	//针对mxn矩阵的最内层：n>m情况下，且最内层是单列
	if a!=c&&b==d {
		for i:=a; i<=c; i++{
			*res=append(*res,matrix[i][d])
		}
		return
	}
	for i:=b; i<=d; i++{
		*res=append(*res, matrix[a][i])
	}
	for i:=a+1; i<=c; i++{
		*res=append(*res,matrix[i][d])
	}
	for i:=d-1; i>=b; i--{
		*res=append(*res,matrix[c][i])
	}
	for i:=c-1; i>a; i--{
		*res=append(*res,matrix[i][b])
	}
}

