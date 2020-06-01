package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	mx := candies[0]
	for i := 0; i < len(candies); i++ {
		mx = max(candies[i], mx)
		candies[i] += extraCandies
	}
	res := make([]bool, len(candies))
	for i := 0; i < len(res); i++ {
		if candies[i] >= mx {
			res[i] = true
		}
	}
	return res
}
