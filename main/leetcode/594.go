package main

//哈希表
//func findLHS(nums []int) int {
//	mp:=make(map[int]int)
//	for _,v:=range nums{
//		mp[v]++
//	}
//	res:=0
//	for k,v:=range mp{
//		if mp[k+1]!=0{
//			res=max(res,v+mp[k+1])
//		}
//	}
//	return res
//}

//排序
//func findLHS(nums []int) int {
//	sort.Ints(nums)
//	mp:=make(map[int]int)
//	mp[nums[0]]=1
//	res:=0
//	nums=append(nums,math.MaxInt32)
//	for k,v:=range nums{
//		if k>0&&v!=nums[k-1]{
//			mp[nums[k]]=k+1
//			if mp[nums[k-1]-1]!=0{
//				res=max(res,k-mp[nums[k-1]-1]+1)
//			}
//		}
//	}
//	return res
//}

func findLHS(nums []int) int {
	mp := make(map[int][]int)
	res := 0
	for _, v := range nums {
		if _, ok := mp[v]; !ok {
			mp[v] = make([]int, 2)
		}
		_, ok1 := mp[v-1]
		if ok1 {
			mp[v-1][0]++
			mp[v][1] = mp[v-1][0]
			res = max(res, mp[v-1][0])
		} else {
			mp[v][1]++
		}
		_, ok2 := mp[v+1]
		if ok2 {
			mp[v+1][1]++
			mp[v][0] = mp[v+1][1]
			res = max(res, mp[v+1][1])
		} else {
			mp[v][0]++
		}
	}
	return res
}
