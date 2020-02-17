package main

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var res [][]int
	if len(nums)==0 {
		return [][]int{{}}
	}
	sort.Ints(nums)
	used:=make([]int,len(nums))
	var path []int
	dfs(nums,0,&path,used,&res)
	return res
}

func dfs(arr []int,depth int,path *[]int,used []int,res *[][]int)  {
	if depth==len(arr){
		//保存数组并终止递归
		temp:=make([]int,len(arr))
		copy(temp,*path)
		*res= append(*res, temp)
		return
	}
	for i:=0; i<len(arr); i++ {
		if used[i]==0 {
			if i>0&&arr[i]==arr[i-1]&&used[i-1]==0 {
				continue
			}
			*path=append(*path, arr[i])
			used[i]=1
			dfs(arr,depth+1,path,used,res)
			//回溯，起始就是used数组标记回溯+path数组减一（切片实现很麻烦）
			used[i]=0
			temp2:=*path
			temp2=temp2[:len(temp2)-1]
			*path= temp2
		}
	}
}