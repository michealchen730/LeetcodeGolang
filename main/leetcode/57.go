package main

func insert(intervals [][]int, newInterval []int) [][]int {
	i := 0
	flag := 0
	for i < len(intervals) {
		if newInterval[0] >= intervals[i][0] {
			if newInterval[0] <= intervals[i][1] {
				flag = 1
				break
			} else {
				i++
			}
		} else {
			break
		}
	}
	j := i
	flag2 := 0
	for j < len(intervals) {
		if newInterval[1] >= intervals[j][0] {
			if newInterval[1] <= intervals[j][1] {
				flag2 = 1
				break
			} else {
				j++
			}
		} else {
			break
		}
	}
	if flag == 0 && flag2 == 0 {
		intervals = append(intervals[:i], append([][]int{newInterval}, intervals[j:]...)...)
	}
	if flag == 1 && flag2 == 1 {
		intervals = append(intervals[:i], append([][]int{{intervals[i][0], intervals[j][1]}}, intervals[j+1:]...)...)
	}
	if flag == 0 && flag2 == 1 {
		newInterval[1] = intervals[j][1]
		intervals = append(intervals[:i], append([][]int{newInterval}, intervals[j+1:]...)...)
	}
	if flag == 1 && flag2 == 0 {
		newInterval[0] = intervals[i][0]
		intervals = append(intervals[:i], append([][]int{newInterval}, intervals[j:]...)...)
	}
	return intervals
}
