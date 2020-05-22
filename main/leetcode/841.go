package main

func canVisitAllRooms(rooms [][]int) bool {
	used := make([]int, len(rooms))
	used[0] = 1
	queue := []int{0}
	for len(queue) != 0 {
		t := queue[0]
		keys := rooms[t]
		for _, v := range keys {
			if used[v] == 0 {
				used[v] = 1
				queue = append(queue, v)
			}
		}
		queue = queue[1:]
	}
	res := 0
	for _, v := range used {
		if v == 0 {
			res++
		}
	}
	return res == len(rooms)
}
