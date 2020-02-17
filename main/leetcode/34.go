package main

func searchRange(nums []int, target int) []int {
	if len(nums)==0 {
		return []int{-1,-1}
	}
	res:=binarySearch2(nums,target,0,len(nums)-1)
	if res==-1{
		return []int{-1,-1}
	}else{
		i,j:=res,res
		for i>=0&&nums[i]==target{
			i--
		}
		for j<=len(nums)-1&&nums[j]==target{
			j++
		}
		return []int{i+1,j-1}
	}
}

func binarySearch2(nums []int,target int, low int, high int) int {
	if target>nums[high]||target<nums[low]||low>high {
		return -1
	}
	middle:=(low+high)/2
	if nums[middle]>target {
		return binarySearch2(nums,target,low,middle-1)
	}else if nums[middle]<target{
		return binarySearch2(nums,target,middle+1,high)
	}else{
		return middle
	}
}
