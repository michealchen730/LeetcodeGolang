package main

func equalSubstring(s string, t string, maxCost int) int {
	start := 0
	end := 0
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += abs(int(s[i]) - int(t[i]))
		end = i
		if sum > maxCost {
			break
		}
	}
	res := end - start

	for end < len(s) {

		for sum > maxCost {
			sum -= abs(int(s[start]) - int(t[start]))
			if start <= end {
				start++
			}
		}

		if sum <= maxCost {
			res = max(end-start+1, res)
		}

		if end == len(s)-1 {
			break
		}

		if start > end {
			sum += abs(int(s[start]) - int(t[start]))
			end++
			continue
		}

		//start==end，说明该下标可用
		if start <= end {
			end++
			sum += abs(int(s[end]) - int(t[end]))
			continue
		}

	}

	return res
}
