package main

func combine(n int, k int) [][]int {
	if k>n {
		return nil
	}
	var res [][]int
	var arr []int
	if n==0 {
		return [][]int{{}}
	}
	getCombine(1,k,&arr,&res,n)
	return res
}

func getCombine(n int,k int,arr *[]int,res *[][]int,max int){
	if len(*arr)==k {
		temp:=make([]int,len(*arr))
		copy(temp,*arr)
		*res=append(*res, temp)
		return
	}
	 if k-len(*arr)>max-n+1{
		 return
	 }
	 //下面屏蔽的代码是剪枝方式，但完全可以用上面的剪枝方式取代，而且上面更优。
	//if n>max {
	//	return
	//}
	getCombine(n+1,k,arr,res,max)
	*arr= append(*arr, n)
	getCombine(n+1,k,arr,res,max)
	res2:=*arr
	res2=res2[:len(res2)-1]
	*arr=res2
}
