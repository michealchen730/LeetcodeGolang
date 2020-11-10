package main

import (
	"math"
)

func totalFruit(tree []int) int {
	if len(tree) < 3 {
		return len(tree)
	}
	res, f1, f2, e1, e2, j := 0, math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32, 0
	for ; j < len(tree); j++ {
		if f1 == math.MaxInt32 {
			f1, e1 = j, j
			continue
		}
		if f2 == math.MaxInt32 && tree[j] != tree[f1] {
			f2, e2 = j, j
			continue
		}
		if tree[j] == tree[f1] || tree[j] == tree[f2] {
			if tree[j] == tree[f1] {
				e1 = j
			} else {
				e2 = j
			}
		} else {
			//这里tree[j]代表出现的第三个值
			//进行清算
			res = max(res, j-min(f2, f1))
			if tree[j-1] == tree[f1] {
				//更新f1
				f1, f2, e2 = e2+1, j, j
			} else {
				f2, f1, e1 = e1+1, j, j
			}
		}
	}
	return max(res, j-min(f2, f1))
}
