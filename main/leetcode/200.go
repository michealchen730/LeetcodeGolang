package main

func numIslands(grid [][]byte) int {
	arr := make([]int, len(grid)*len(grid[0]))
	for i, _ := range arr {
		arr[i] = i
	}
	res := 0
	r := len(grid)
	c := len(grid[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				res++
				if checkI(grid, i-1, j) {
					res -= unionIslands(arr, (i-1)*c+j, i*c+j)
				}
				if checkI(grid, i, j-1) {
					res -= unionIslands(arr, i*c+j-1, i*c+j)
				}
			}
		}
	}
	return res
}

func findRootIsland(arr []int, i int) int {
	for arr[i] != i {
		i = arr[i]
	}
	return i
}

func unionIslands(arr []int, i, j int) int {
	m, n := findRootIsland(arr, i), findRootIsland(arr, j)
	if m != n {
		arr[m] = arr[n]
		return 1
	}
	return 0
}

func checkI(grid [][]byte, i, j int) bool {
	if i >= 0 && i <= len(grid)-1 && j >= 0 && j <= len(grid[0])-1 && grid[i][j] == '1' {
		return true
	}
	return false
}

//func numIslands(grid [][]byte) int {
//	res:=0
//	for i:=0;i<len(grid);i++{
//		for j:=0;j<len(grid[0]);j++{
//			if grid[i][j]=='1' {
//				res++
//				stack:=[][]int{{i,j}}
//				for len(stack)!=0{
//					temp:=stack[0]
//					grid[temp[0]][temp[1]]='U'
//					if checkI(grid,temp[0]-1,temp[1]) {
//						grid[temp[0]-1][temp[1]]='U'
//						stack=append(stack, []int{temp[0]-1,temp[1]})
//					}
//					if checkI(grid,temp[0]+1,temp[1]) {
//						grid[temp[0]+1][temp[1]]='U'
//						stack=append(stack, []int{temp[0]+1,temp[1]})
//					}
//					if checkI(grid,temp[0],temp[1]-1) {
//						grid[temp[0]][temp[1]-1]='U'
//						stack=append(stack, []int{temp[0],temp[1]-1})
//					}
//					if checkI(grid,temp[0],temp[1]+1) {
//						grid[temp[0]][temp[1]+1]='U'
//						stack=append(stack, []int{temp[0],temp[1]+1})
//					}
//					stack=stack[1:]
//				}
//			}
//		}
//	}
//	return res
//}

//func checkI(grid [][]byte, i,j int) bool {
//	if i>=0&&i<=len(grid)-1&&j>=0&&j<=len(grid[0])-1&&grid[i][j]=='1' {
//		return true
//	}
//	return false
//}
