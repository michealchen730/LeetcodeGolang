package main

func openLock(deadends []string, target string) int {
	deadmap := make(map[string]bool)
	for _, v := range deadends {
		deadmap[v] = true
	}

	queue := []string{"0000"}
	depth := 0
	for len(queue) != 0 {
		length := len(queue)
		for idx := 0; idx < length; idx++ {
			if deadmap[queue[idx]] {
				continue
			} else {
				deadmap[queue[idx]] = true
			}
			if queue[idx] == target {
				return depth
			}
			paths := nextStep(queue[idx])
			queue = append(queue, paths...)
		}
		queue = queue[length:]
		depth += 1
	}
	return -1
}

func nextStep(l string) []string {
	bytes := []byte(l)
	var res []string
	for k, v := range bytes {
		bytes[k] = v + 1
		if v == '9' {
			bytes[k] = '0'
		}
		res = append(res, string(bytes))
		bytes[k] = v - 1
		if v == '0' {
			bytes[k] = '9'
		}
		res = append(res, string(bytes))
		bytes[k] = v
	}
	return res
}
