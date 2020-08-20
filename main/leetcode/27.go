package main

//还是双指针
//题意允许元素的顺序可以改变，但这里的解法不改变元素的顺序
func removeElement(nums []int, val int) int {
	left := 0
	for left < len(nums) {
		if nums[left] != val {
			left++
		} else {
			right := left + 1
			for ; right < len(nums); right++ {
				if nums[right] != val {
					nums[left], nums[right] = nums[right], nums[left]
					left++
					break
				}
			}
			if right == len(nums) {
				break
			}
		}
	}
	return left
}
