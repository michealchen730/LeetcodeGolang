package main

func permute(nums []int) [][]int {
	var res [][]int
	toNext(nums,0,&res)
	return res
}
func toNext(arr []int,n int,res *[][]int)  {
	if n==len(arr)-1 {
		temp:=make([]int,len(arr))
		copy(temp,arr)
		*res=append(*res, temp)
	}
	for i:=n; i<len(arr); i++ {
		arr[i],arr[n]=arr[n],arr[i]
		toNext(arr,n+1,res)
		arr[i],arr[n]=arr[n],arr[i]
	}
}