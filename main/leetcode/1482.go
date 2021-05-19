package main

//func minDays(bloomDay []int, m int, k int) int {
//	least:=m*k
//	n:=len(bloomDay)
//	if least<=n{
//		cpy:=make([]int,n)
//		copy(cpy,bloomDay)
//		sort.Ints(cpy)
//		binArr:=make([]int,0,n-least+1)
//		for t:=least-1;t<n;t++{
//			if t==least-1||cpy[t]!=cpy[t-1]{
//				binArr=append(binArr,cpy[t])
//			}
//		}
//		low,high:=0,len(binArr)-1
//		res:=binArr[high]
//		for low<=high{
//			mid:=low+(high-low)/2
//			if canSolve1482(bloomDay,m,k,binArr[mid]){
//				res=binArr[mid]
//				high=mid-1
//			}else{
//				low=mid+1
//			}
//		}
//		return res
//	}
//	return -1
//}

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	maxD := 0
	for _, v := range bloomDay {
		if v > maxD {
			maxD = v
		}
	}
	low, high, res := 0, maxD, maxD
	for low <= high {
		mid := low + (high-low)/2
		if canSolve1482(bloomDay, m, k, mid) {
			res = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}

func canSolve1482(bloomDay []int, m, k, l int) bool {
	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] <= l && len(bloomDay)-i >= k {
			flag := true
			for j := i + 1; j < i+k; j++ {
				if bloomDay[j] > l {
					i = j
					flag = false
					break
				}
			}
			if flag {
				i = i + k - 1
				m--
				if m <= 0 {
					return true
				}
			}
		}
	}
	return false
}
