package main

func findBall(grid [][]int) []int {
	res := make([]int, len(grid[0]))
	for k := range res {
		res[k] = k
	}
	for _, v := range grid {
		for k, r := range res {
			if r != -1 {
				//每个球在找下一轮的所在位置时，一定要考虑当前位置和它的邻居位置
				//找左邻居还是右邻居，要看当前的挡板方向，是1，要去找右边的，是0则去找左边的
				//如果已经到达边界，不存在左右位置的挡板了，直接置-1
				//如果邻居的挡板方向和自己的挡板方向一致，那么相应的加减1即可
				if v[r] == 1 && r+1 < len(v) && v[r+1] == 1 {
					res[k] += 1
				} else if v[r] == -1 && r > 0 && v[r-1] == -1 {
					res[k] -= 1
				} else {
					res[k] = -1
				}
			}
		}
	}
	return res
}
