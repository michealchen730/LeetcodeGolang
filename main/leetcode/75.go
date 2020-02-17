package main

func sortColors(nums []int)  {
	if len(nums)==0 {
		return
	}
	//碰见2就往后放
	//碰见0就往前放
	i,j:=0,len(nums)-1
	temp:=0
	for temp<len(nums){
		if nums[temp]==2&&temp<j{
			if nums[j]!=2 {
				nums[j],nums[temp]=nums[temp],nums[j]
			}else{
				j--
			}
		}else if nums[temp]==0&&temp>i{
			if nums[i]!=0 {
				nums[i],nums[temp]=nums[temp],nums[i]
			}else{
				i++
			}
		}else {
			temp++
		}
	}
}