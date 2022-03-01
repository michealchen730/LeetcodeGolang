package main

import "strings"

//最简单的思路是模拟
//当前施加力的骨牌一定会倒，并且会影响到左右邻居中的一个（根据方向）
//对每一轮进行结算，然后循环往复即可

//第二种思路是直接结算
//找到所有施加力的骨牌，按照方向结算
func pushDominoes(dominoes string) string {
	var res strings.Builder
	tempLength := 0
	Right := '.'
	for k, v := range dominoes {
		//出现R,R之前的要结算
		if v == 'R' {
			for tempLength < k {
				res.WriteRune(Right)
				tempLength++
			}
			Right = 'R'
		}
		if v == 'L' {
			if Right == '.' {
				for tempLength <= k {
					res.WriteRune('L')
					tempLength++
				}
			} else {
				t := (k - tempLength + 1) / 2
				for i := 0; i < t; i++ {
					res.WriteRune('R')
				}
				if (k-tempLength)%2 == 0 {
					res.WriteRune('.')
				}
				for i := 0; i < t; i++ {
					res.WriteRune('L')
				}
				tempLength = k + 1
			}
			Right = '.'
		}
	}
	for tempLength < len(dominoes) {
		res.WriteRune(Right)
		tempLength++
	}
	return res.String()
}
