package main

func canJump(nums []int) bool {
	if len(nums)==1{
		return true
	}
	return goBack(nums,len(nums)-1)
}

func goBack(nums []int,length int)  bool{
	for i:=0; i<=length; i++ {
		//无解，终止迭代
		if i==length {
			return false
		}
		if i+nums[i]>=length {
			//找到解，终止迭代
			if i==0 {
				return true
			}
			return goBack(nums,i)
		}
	}
	return false
}


