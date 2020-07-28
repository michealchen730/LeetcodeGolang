package main

func twoSum167(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		t := numbers[i] + numbers[j]
		if t == target {
			return []int{i + 1, j + 1}
		} else if t > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}
