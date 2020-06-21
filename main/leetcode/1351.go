package main

//数据可以暴力，但可能想考察的还是二分
func countNegatives(grid [][]int) int {
	res := 0
	low, high := 0, len(grid)-1
	s := high
	//只有一行
	if high == 0 {
		s = high + 1
		if len(grid[0]) == 0 {
			if grid[0][0] < 0 {
				return 1
			} else {
				return 0
			}
		}
	} else {
		//找到某行的最开头是小于0的
		for low < high {
			i := low + (high-low)/2
			j := i + 1
			if grid[i][0] >= 0 && grid[j][0] < 0 {
				res += (len(grid) - j) * len(grid[0])
				s = j
				break
			}
			if grid[j][0] >= 0 {
				low = j
			}
			if grid[i][0] < 0 {
				high = i
			}
		}
		if low == high {
			//没找到，两种情况，所有行开头都小于0
			if low == 0 {
				res += len(grid) * len(grid[0])
				return res
			} else {
				//所有行开头均大于0
				s = len(grid)
			}
		}
		//如果只有一列
		if len(grid[0]) == 1 {
			return res
		}
	}
	//不止一列
	for _, v := range grid[0:s] {
		low, high := 0, len(v)-1
		for low < high {
			i := low + (high-low)/2
			j := i + 1
			if v[i] >= 0 && v[j] < 0 {
				res += len(v) - j
				break
			}
			if v[j] >= 0 {
				low = j
			}
			if v[i] < 0 {
				high = i
			}
		}
		if low == high && low == 0 {
			res += len(grid[0])
		}
	}
	return res
}
