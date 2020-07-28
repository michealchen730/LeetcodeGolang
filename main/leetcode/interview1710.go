package main

func majorityElementI1710(nums []int) int {
	sum := 0
	temp := 0
	for _, v := range nums {
		if sum == 0 {
			sum++
			temp = v
			continue
		}
		if v != temp {
			sum--
		} else {
			sum++
		}
	}
	count := 0
	for _, v := range nums {
		if v == temp {
			count++
		}
	}
	if count > len(nums)/2 {
		return temp
	}
	return -1
}
