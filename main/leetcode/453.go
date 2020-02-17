package main

func minMoves(nums []int) int {
	if len(nums)==1 {
		return 0
	}
	sum:=0
	temp:=nums[0]
	for i:=1; i<len(nums); i++ {
		if nums[i]<temp {
			temp=nums[i]
		}
	}
	for i:=0; i<len(nums); i++{
		if nums[i]>temp {
			sum=sum+nums[i]-temp
		}
	}
	return sum
}
