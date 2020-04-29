package main

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	arr := make([][]int, 8)
	used := make([]int, 8)
	for k, _ := range arr {
		arr[k] = make([]int, 2)
	}
	for _, v := range queens {
		if v[0] == king[0] {
			if v[1] < king[1] {
				if used[0] == 0 || arr[0][1] < v[1] {
					used[0] = 1
					arr[0] = v
				}
			} else {
				if used[1] == 0 || arr[1][1] > v[1] {
					used[1] = 1
					arr[1] = v
				}
			}
		}
		if v[1] == king[1] {
			if v[0] < king[0] {
				if used[2] == 0 || arr[2][0] < v[0] {
					used[2] = 1
					arr[2] = v
				}
			} else {
				if used[3] == 0 || arr[3][0] > v[0] {
					used[3] = 1
					arr[3] = v
				}
			}
		}
		if v[0]-king[0] == v[1]-king[1] {
			if v[0] > king[0] {
				if used[4] == 0 || arr[4][0] > v[0] {
					used[4] = 1
					arr[4] = v
				}
			} else {
				if used[5] == 0 || arr[5][0] < v[0] {
					used[5] = 1
					arr[5] = v
				}
			}
		}
		if v[0]-king[0] == king[1]-v[1] {
			if v[0] > king[0] {
				if used[6] == 0 || arr[6][0] > v[0] {
					used[6] = 1
					arr[6] = v
				}
			} else {
				if used[7] == 0 || arr[7][0] < v[0] {
					used[7] = 1
					arr[7] = v
				}
			}
		}
	}
	var res [][]int
	for k, v := range used {
		if v == 1 {
			res = append(res, arr[k])
		}
	}
	return res
}
