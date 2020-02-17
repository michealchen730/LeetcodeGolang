package main

import (
	"math"
)


func solveNQueens(n int) [][]string {
	var res [][]string
	chessboard:=make([]int,n)
	placeQueen_1(&res,chessboard,n,0)
	return res
}

func checkResult_1(chessboard []int,n int) bool {
	for i:=0; i<n; i++ {
		if chessboard[i]==chessboard[n]||math.Abs(float64(chessboard[i]-chessboard[n]))==float64(n-i){
			return false
		}
	}
	return true
}

func placeQueen_1(result *[][]string,chessboard []int,max int,temp int)  {
	if temp==max {
		output(chessboard,result)
		return
	}
	for i:=0; i<max; i++ {
		chessboard[temp]=i
		if checkResult(chessboard,temp) {
			placeQueen_1(result,chessboard,max,temp+1)
		}
	}
}

func output(chessboard []int,result *[][]string) {
	var tempRes []string
	for _,v:=range chessboard{
		ts:=""
		for i:=0; i<len(chessboard); i++ {
			if i==v{
				ts=ts+"Q"
			}else{
				ts=ts+"."
			}
		}
		tempRes = append(tempRes, ts)
	}
	*result=append(*result, tempRes)
}
