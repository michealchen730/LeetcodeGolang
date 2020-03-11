package main

//这题重做。。。
//func dailyTemperatures(T []int) []int {
//	var res []int
//	if len(T)==0{
//		return res
//	}
//	res=append(res, 0)
//	if len(T)==1{
//		return res
//	}
//	for i:=len(T)-2;i>=0;i--{
//		if T[i]<T[i+1] {
//			res=append([]int{1},res...)
//		}else{
//			temp:=res[0]
//			for T[i+1+temp]<=T[i] {
//				if res[temp]==0 {
//					res=append([]int{0},res...)
//					break
//				}else{
//					temp+=res[temp]
//				}
//			}
//			if T[i+1+temp]>T[i] {
//				res=append([]int{1+temp},res...)
//			}
//		}
//	}
//	return res
//}

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	var stack []int
	i := len(T) - 1
	for i >= 0 {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			for true {
				if len(stack) != 0 && T[i] >= T[stack[len(stack)-1]] {
					stack = stack[:len(stack)-1]
				} else {
					if len(stack) != 0 {
						res[i] = stack[len(stack)-1] - i
					}
					stack = append(stack, i)
					break
				}
			}
		}
		i--
	}
	return res
}
