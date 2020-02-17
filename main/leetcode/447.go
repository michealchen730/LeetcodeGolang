package main

func numberOfBoomerangs(points [][]int) int {
	if len(points)<3 {
		return 0
	}
	matrix:=make([][]int,len(points))
	for i:=0;i<len(points);i++ {
		matrix[i]=make([]int,len(points))
	}
	for i:=0; i<len(points)-1; i++ {
		for j:=i+1;j<len(points);j++{
			matrix[i][j]=(points[i][0]-points[j][0])*(points[i][0]-points[j][0])+(points[i][1]-points[j][1])*(points[i][1]-points[j][1])
			matrix[j][i]=matrix[i][j]
		}
	}
	sum:=0
	for i:=0; i< len(matrix); i++ {
		tempMap:=make(map[int]int)
		for j:=0; j<len(matrix); j++ {
			if matrix[i][j]!=0 {
				tempMap[matrix[i][j]]++
			}
		}
		for _,v:=range tempMap{
			println(v)
			if v>=2 {
				sum=sum+v*(v-1)
			}
		}
	}
	return sum
}

