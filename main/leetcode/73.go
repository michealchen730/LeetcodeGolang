package main

func setZeroes(matrix [][]int)  {
	if len(matrix)==0 {
		return
	}
	col:=make(map[int]int)
	row:=make(map[int]int)
	for i:=0; i<len(matrix); i++ {
		for j:=0; j<len(matrix[0]); j++ {
			if matrix[i][j]==0 {
				col[i]=1
				row[j]=1
			}
		}
	}
	for k,_:=range col {
		for i:=0; i<len(matrix[0]); i++ {
			matrix[k][i]=0
		}
	}
	for k,_:=range row {
		for i:=0; i<len(matrix); i++ {
			matrix[i][k]=0
		}
	}
}
