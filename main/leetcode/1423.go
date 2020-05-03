package main

func maxScore2(cardPoints []int, k int) int {
	var res []int
	for i := k - 1; i >= 0; i-- {
		res = append(res, cardPoints[i])
	}
	for i := len(cardPoints) - 1; i >= len(cardPoints)-k; i-- {
		res = append(res, cardPoints[i])
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += res[i]
	}
	mx := sum
	for i := k; i < len(res); i++ {
		sum -= res[i-k]
		sum += res[i]
		if sum > mx {
			mx = sum
		}
	}
	return mx
}
