package main

//暴力法
//超时
//func wordSubsets(A []string, B []string) []string {
//	var res []string
//	flag:=make([]bool,len(A))
//	//对每一个A中的元素
//	for idx,v:=range A{
//		arr:=make([]int,26)
//		for k,_:=range v{
//			arr[v[k]-'a']++
//		}
//		tmp:=make([]int,26)
//		copy(tmp,arr)
//		//都要去比对每一个B中的元素
//		SE:
//		for _,val:=range B{
//			for i,_:=range val{
//				arr[val[i]-'a']--
//				//一旦不满足，flag数组直接置位
//				if arr[val[i]-'a']<0{
//					flag[idx]=true
//					break SE
//				}
//			}
//			copy(arr,tmp)
//		}
//	}
//	for k,v:=range flag{
//		if !v{
//			res=append(res,A[k])
//		}
//	}
//	return res
//
//}

//题解思路，将B合成一个单词
//很厉害，也是最开始没看清题意，所以越想越歪
func wordSubsets(A []string, B []string) []string {
	arr := make([]int, 26)
	for _, v := range B {
		tmp := make([]int, 26)
		for k, _ := range v {
			tmp[v[k]-'a']++
		}
		for k, val := range tmp {
			if val > arr[k] {
				arr[k] = val
			}
		}
	}
	var res []string
	for _, v := range A {
		tmp := make([]int, 26)
		for k, _ := range v {
			tmp[v[k]-'a']++
		}
		flag := true
		for k, val := range tmp {
			if val < arr[k] {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, v)
		}
	}
	return res
}
