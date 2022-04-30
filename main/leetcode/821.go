package main

//BFS
//func shortestToChar(s string, c byte) []int {
//	res:=make([]int,len(s))
//	for k,_:=range res{
//		res[k]=-1
//	}
//
//	var queue []int
//
//	for k,_:=range s{
//		if s[k]==c{
//			res[k]=0
//			queue=append(queue,k)
//		}
//	}
//	tmpDistance:=1
//	for len(queue)!=0{
//		length:=len(queue)
//		for i:=0;i<length;i++{
//			if queue[i]+1<len(res)&&res[queue[i]+1]==-1{
//				res[queue[i]+1]=tmpDistance
//				queue=append(queue,queue[i]+1)
//			}
//			if queue[i]-1>=0&&res[queue[i]-1]==-1{
//				res[queue[i]-1]=tmpDistance
//				queue=append(queue,queue[i]-1)
//			}
//		}
//		queue=queue[length:]
//		tmpDistance++
//	}
//	return res
//}

//两次遍历（官方思路）
func shortestToChar(s string, c byte) []int {
	tmpIndexOfC := len(s) - 1
	res := make([]int, len(s))
	for k, _ := range s {
		if c == s[k] {
			tmpIndexOfC = k
		} else {
			res[k] = abs(k - tmpIndexOfC)
		}
	}

	tmpIndexOfC = 0
	for i := len(s) - 1; i >= 0; i-- {
		if c == s[i] {
			tmpIndexOfC = i
		} else {
			res[i] = min(res[i], abs(i-tmpIndexOfC))
		}
	}
	return res
}
