package main

func maxDistToClosest(seats []int) int {
	if len(seats) == 1 {
		return 1
	}
	dist1, first, last1, maxdist := 0, -1, -1, 0
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if first == -1 {
				dist1, first, last1 = i-0, 0, i
			}
			if last1 != -1 {
				if maxdist < i-last1 {
					maxdist = i - last1
				}
			}
			last1 = i
		}
	}
	if len(seats)-1-last1 > dist1 {
		dist1 = len(seats) - 1 - last1
	}
	if maxdist/2 < dist1 {
		return dist1
	} else {
		return maxdist / 2
	}
}
