package main

func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	c1 := checkClimax(nums, 0, len(nums)-1)
	if c1 == len(nums)-1 || c1 == len(nums)-2 {
		return true
	}
	if c1 == 0 {
		return checkClimax(nums, 1, len(nums)-1) == len(nums)-1
	}
	temp := nums[c1]
	nums[c1] = nums[c1-1]
	bool1 := checkClimax(nums, c1, len(nums)-1) == len(nums)-1
	nums[c1] = temp
	nums[c1+1] = nums[c1+2]
	bool2 := checkClimax(nums, c1, len(nums)-1) == len(nums)-1
	return bool1 || bool2
}

/**
如果返回的i==high，说明是个非递减序列
*/
func checkClimax(nums []int, low int, high int) int {
	i := low
	for i < high && nums[i] <= nums[i+1] {
		i++
	}
	return i
}

//if len(nums)<=2 {
//	return true
//}
//p1:=0
//for p1+1<len(nums)&&nums[p1]<=nums[p1+1]{
//	p1++
//}
//if p1==len(nums)-1 {
//	return true
//}else {
//	p2:=p1+1
//	if p2==len(nums)-1 {
//		return true
//	}
//	if p1-1<0||nums[p2]>=nums[p1-1] {
//		p2=p2+1
//		for p2<len(nums)&&nums[p2]>=nums[p2-1]{
//			p2++
//		}
//		if p2==len(nums) {
//			return true
//		}else{
//			return false
//		}
//	}else {
//		return false
//	}
//}
//}
