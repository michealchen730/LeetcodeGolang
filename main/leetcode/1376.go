package main

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	arr := make([]int, len(manager))
	max := 0
	for i := 0; i < len(arr); i++ {
		time := 0
		t := i
		for manager[t] != -1 {
			if manager[t] < i {
				time += arr[manager[t]] + informTime[manager[t]]
				break
			}
			time += informTime[manager[t]]
			t = manager[t]
		}
		arr[i] = time
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
