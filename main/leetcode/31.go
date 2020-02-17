package main

func nextPermutation(nums []int)  {
	if len(nums)<=1 {
		return
	}
	i:=len(nums)-1

	for i-1>=0&&nums[i-1]>=nums[i] {
		i--
	}
	if i==0 {
		reverse(nums)
	}else{
		for temp:=len(nums)-1;temp>=0;temp--{
			if nums[temp]>nums[i-1] {
				nums[temp],nums[i-1]=nums[i-1],nums[temp]
				reverse(nums[i:len(nums)])
				break
			}
		}
	}
}
func reverse(arr []int)  {
	i,j:=0,len(arr)-1
	for i<j {
		arr[i],arr[j]=arr[j],arr[i]
		i++
		j--
	}
}
