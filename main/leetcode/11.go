package main

//从两边到中间
//挑两边最小的那个开始移动
func maxArea(height []int) int {
	i,j:=0,len(height)-1
	res:=min(height[i],height[j])*(j-i)
	for i<j{
		temp:=0
		if height[i]<height[j] {
			temp=height[i]
			for height[i]<=temp&&i<j {
				i++
			}
			if min(height[i],height[j])*(j-i)>res {
				res=min(height[i],height[j])*(j-i)
			}
		}else {
			temp=height[j]
			for height[j]<=temp&&i<j{
				j--
			}
			if min(height[i],height[j])*(j-i)>res {
				res=min(height[i],height[j])*(j-i)
			}
		}
	}
	return res
}

func min(x int,y int) int {
	if x>y {
		return y
	}else {
		return x
	}
}
