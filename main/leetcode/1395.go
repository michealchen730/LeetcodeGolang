package main

func numTeams(rating []int) int {
	res := 0
	for i := 0; i < len(rating)-2; i++ {
		for j := i + 1; j < len(rating)-1; j++ {
			for k := j + 1; k < len(rating); k++ {
				if (rating[i] < rating[j] && rating[j] < rating[k]) || (rating[i] > rating[j] && rating[j] > rating[k]) {
					res++
				}
			}
		}
	}
	return res
}
