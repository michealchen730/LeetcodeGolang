package main

func minOperations1713(target []int, arr []int) int {
	mp := make(map[int]int)
	for k, v := range target {
		mp[v] = k
	}
	var arr2 []int
	for _, v := range arr {
		if idx, ok := mp[v]; ok {
			arr2 = append(arr2, idx)
		}
	}
	var tail []int
	for _, v := range arr2 {
		if len(tail) == 0 || v > tail[len(tail)-1] {
			tail = append(tail, v)
		} else {
			low, high := 0, len(tail)-1
			flag := false
			for low < high {
				mid := low + (high-low)/2
				if tail[mid] > v {
					high = mid - 1
				} else if tail[mid] < v {
					low = mid + 1
				} else {
					flag = true
					break
				}
			}
			if flag {
				continue
			} else {
				if tail[low] < v {
					tail[low+1] = v
				} else {
					tail[low] = v
				}
			}
		}
	}
	return len(target) - len(tail)
}
