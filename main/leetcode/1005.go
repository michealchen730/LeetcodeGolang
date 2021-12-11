package main

import (
	"math"
	"sort"
)

//func largestSumAfterKNegations(A []int, K int) int {
//	sort.Ints(A)
//	start:=0
//	minNeg:=math.MinInt32
//	for K>0{
//		if start>=len(A){
//			break
//		}
//		if A[start]<0{
//			if A[start]>minNeg{
//				minNeg=-A[start]
//			}
//			A[start]=-A[start]
//			start++
//			K--
//		}else{
//			break
//		}
//	}
//	sum:=0
//	for _,v:=range A{
//		sum+=v
//	}
//	if K%2==0{
//		return sum
//	}else{
//		if start<len(A){
//			minNeg=min(minNeg,A[start])
//		}
//		return sum-2*minNeg
//	}
//}

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	start := 0
	minNeg := math.MaxInt32
	sum := 0
	for K > 0 && start < len(A) && A[start] < 0 {
		sum += -A[start]
		if -A[start] < minNeg {
			minNeg = -A[start]
		}
		start++
		K--
	}
	for start < len(A) {
		minNeg = min(minNeg, A[start])
		sum += A[start]
		start++
	}
	if K == 0 || K%2 == 0 {
		return sum
	} else {
		return sum - 2*minNeg
	}
}
