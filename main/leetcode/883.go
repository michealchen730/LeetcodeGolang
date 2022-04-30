package main

//func projectionArea(grid [][]int) int {
//	res:=0
//	row:=make([]int,len(grid))
//	for idx,line := range grid{
//		for col,val := range line{
//			if val>0{
//				res+=1
//			}
//			row[idx]=max(row[idx],val)
//			grid[0][col]=max(grid[0][col],val)
//		}
//	}
//	for _,val:=range grid[0]{
//		res+=val
//	}
//	for _,val:=range row{
//		res+=val
//	}
//	return res
//}

//相比上面的代码，应该可以少用一行空间
func projectionArea(grid [][]int) int {
	res := 0
	r1 := 0
	for idx, line := range grid {
		for col, val := range line {
			if val > 0 {
				res += 1
			}
			grid[0][col] = max(grid[0][col], val)
			if idx == 0 {
				r1 = max(r1, val)
			} else {
				grid[idx][0] = max(grid[idx][0], val)
			}
		}
		if idx != 0 {
			res += grid[idx][0]
		}
	}
	for _, val := range grid[0] {
		res += val
	}
	res = res + r1
	return res
}
