package main

import (
	"math"
)

var totalRes int
func totalNQueens(n int) int {
	totalRes=0
	chessboard:=make([]int,n)
	placeQueen(chessboard,n,0)
	return totalRes
}

func checkResult(chessboard []int,n int) bool {
	for i:=0; i<n; i++ {
		if chessboard[i]==chessboard[n]||math.Abs(float64(chessboard[i]-chessboard[n]))==float64(n-i){
			return false
		}
	}
	return true
}

func placeQueen(chessboard []int,max int,temp int)  {
	if temp==max {
		totalRes++
		return
	}
	for i:=0; i<max; i++ {
		chessboard[temp]=i
		if checkResult(chessboard,temp) {
			placeQueen(chessboard,max,temp+1)
		}
	}
}

