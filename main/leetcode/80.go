package main

//函数名被我更改了
func removeDuplicates80(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}
	point1 := 0
	point2 := 1
	write := 0
	for point2 != len(nums) {
		if nums[point2] == nums[point1] {
			point2++
		} else {
			nums[write] = nums[point1]
			write++
			if point2-point1 >= 2 {
				nums[write] = nums[point1]
				write++
				length = length - (point2 - point1) + 2
			}
			point1 = point2
			point2++
		}
	}
	nums[write] = nums[point1]
	if point2-point1 >= 2 {
		write++
		nums[write] = nums[point1]
		length = length - (point2 - point1) + 2
	}
	return length
}
