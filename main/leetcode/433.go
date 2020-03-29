package main

func minMutation(start string, end string, bank []string) int {
	used := make([]int, len(bank))
	stack := []string{start}
	nums := 0
	for len(stack) != 0 {
		nums++
		level := len(stack)
		for i := 0; i < level; i++ {
			for k, v := range bank {
				if v != start && used[k] == 0 {
					diff := 0
					for m := 0; m < 8; m++ {
						if v[m] != stack[i][m] {
							diff++
						}
						if diff >= 2 {
							break
						}
					}
					if diff == 1 {
						used[k] = nums
						if bank[k] == end {
							return nums
						}
						stack = append(stack, bank[k])
					}
				}
			}
		}
		stack = stack[level:]
	}
	return -1
}
