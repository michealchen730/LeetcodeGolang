package main

import (
	"math"
)

//1.所有的孩子必须人手一颗
//2.最少的孩子可以只给一颗，然后从它开始扩散

//第2点更正，所有谷底孩子均只给一颗
//分数高的孩子分配到的糖果数目要由他相邻的分数少的孩子来决定

//[2 1 1]这个例子
func candy(ratings []int) int {
	l := len(ratings)
	ratings = append(ratings, math.MaxInt32)
	newRat := append([]int{math.MaxInt32}, ratings...)

	candies := make([]int, len(newRat))
	used := make([]bool, len(newRat))
	gave := 0

	sum := 0

	for gave < l {
		for i := 1; i <= l; i++ {
			if !used[i] {
				//谷底元素赋值1
				if newRat[i] <= newRat[i-1] && newRat[i] <= newRat[i+1] {
					candies[i] = 1
					sum += 1
					used[i] = true
					gave++
				} else if newRat[i] > newRat[i-1] && newRat[i] > newRat[i+1] {
					if used[i+1] && used[i-1] {
						candies[i] = max(candies[i-1], candies[i+1]) + 1
						sum += candies[i]
						used[i] = true
						gave++
					}
				} else if newRat[i] > newRat[i-1] && newRat[i] <= newRat[i+1] {
					if used[i-1] {
						used[i] = true
						candies[i] = candies[i-1] + 1
						sum += candies[i]
						gave++
					}
				} else {
					if used[i+1] {
						used[i] = true
						candies[i] = candies[i+1] + 1
						sum += candies[i]
						gave++
					}
				}
			}
		}
	}
	return sum
}
