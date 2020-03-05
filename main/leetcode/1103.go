package main

func distributeCandies(candies int, num_people int) []int {
	if num_people == 1 {
		return []int{candies}
	}
	arr := make([]int, num_people)
	temp := 1
	for candies > 0 {
		if candies < temp {
			arr[(temp-1)%num_people] += candies
			break
		}
		arr[(temp-1)%num_people] += temp
		candies -= temp
		temp++
	}
	return arr
}
