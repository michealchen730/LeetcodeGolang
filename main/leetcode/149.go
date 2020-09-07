package main

import (
	"strconv"
	"strings"
)

//高精度除法
func Slope(a, b int) string {
	var sb strings.Builder
	var by byte
	sb.WriteString(strconv.Itoa(a / b))
	sb.WriteByte('.')
	for i := 0; i < 20; i++ {
		a = a % b * 10
		by = (byte)(a/b + 48)
		sb.WriteByte(by)
	}
	return sb.String()
}

func maxPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}

	//判定多个点在一条线上，需要
	//1. 任意两点斜率相同
	//2. 某个点的坐标

	//判断某一个点的斜率是否被使用过
	mp2 := make(map[string]map[int]bool)
	mp3 := make(map[string]map[int]int)
	//给定水平线的哈希表
	mp1 := make(map[int]int)
	res := 0
	for i := 0; i < len(points); i++ {
		mp1[points[i][0]]++
		res = max(res, mp1[points[i][0]])
		self := 0
		for j := i + 1; j < len(points); j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				self++
			}

			//先求斜率
			//防止水平线的存在
			if points[i][0] != points[j][0] {
				slope := Slope(points[i][1]-points[j][1], points[i][0]-points[j][0])
				//如果这个点的斜率没有被使用过
				if _, ok := mp2[slope][i]; !ok {
					mp2[slope] = make(map[int]bool)
					mp3[slope] = make(map[int]int)
					mp2[slope][i] = true
					mp3[slope][i]++
					res = max(res, mp3[slope][i])
				}
				//在执行到这一步的时候，该斜率的i点肯定被使用过了
				if !mp2[slope][j] {
					mp2[slope][j] = true
					//这里注意是i不是j，可以自己模拟下三个在一条线上的点测试下算法执行过程
					mp3[slope][i]++
					res = max(res, mp3[slope][i])
				}
			}
		}
		for k, _ := range mp3 {
			mp3[k][i] += self
			res = max(res, mp3[k][i])
		}

	}
	return res
}
