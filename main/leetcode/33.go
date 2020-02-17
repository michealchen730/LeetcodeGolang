package main

//变种的二分查找？
func search(nums []int, target int) int {
	if len(nums)==0 {
		return -1
	}
	if len(nums)==1{
		if target==nums[0]{
			return 0
		}else {
			return -1
		}
	}
	temp:=pointChanged(nums)
	if temp==0 {
		temp=len(nums)-1
		return binarySearch(nums,target,0,temp)
	}
	if target>=nums[0] {
		return binarySearch(nums,target,0,temp-1)
	}else{
		return binarySearch(nums,target,temp,len(nums)-1)
	}
}

func binarySearch(nums []int,target int, low int, high int) int{
	if target>nums[high]||target<nums[low]||low>high {
		return -1
	}
	middle:=(low+high)/2
	if nums[middle]>target {
		return binarySearch(nums,target,low,middle-1)
	}else if nums[middle]<target{
		return binarySearch(nums,target,middle+1,high)
	}else{
		return middle
	}
}

func pointChanged(nums []int) int {
	i,j:=0,len(nums)-1
	if nums[i]<=nums[j] {
		return 0
	}
	//考虑到时间复杂度，我们未必需要在一次遍历里找到，可以先二分找到旋转的点，再去找target
	for i<j{
		temp:=i+(j-i)/2
		//这一步是正常操作
		if nums[temp]>=nums[i]{
			i=temp
		}else{
			j=temp
		}
		if j-i==1{
			if nums[j]<nums[i]{
				return j
			}else{
				return -1
			}
		}
	}
	return j
}
