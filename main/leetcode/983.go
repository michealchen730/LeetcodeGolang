package main

func mincostTickets(days []int, costs []int) int {
	year := make([]int, days[len(days)-1]+1)
	for k, v := range days {
		if k != 0 {
			for i := days[k-1] + 1; i < v; i++ {
				year[i] = year[days[k-1]]
			}
		}
		temp1 := min(costs[0]+year[max(v-1, 0)], costs[1]+year[max(v-7, 0)])
		year[v] = min(temp1, costs[2]+year[max(v-30, 0)])
	}
	return year[days[len(days)-1]]
}
