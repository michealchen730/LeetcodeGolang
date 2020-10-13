package main

func corpFlightBookings(bookings [][]int, n int) []int {
	arr := make([]int, n+2)

	for _, v := range bookings {
		updateRange(v[0], v[1], v[2], arr)
	}

	var res []int

	for i := 1; i <= n; i++ {
		res = append(res, find(i, arr))
	}
	return res
}
