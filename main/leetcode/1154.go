package main

import (
	"strconv"
)

func dayOfYear(date string) int {
	doy := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	sum := 0
	year, _ := strconv.Atoi(date[0:4])
	if (year%100 == 0 && year%400 == 0) || (year%100 != 0 && year%4 == 0) {
		doy[2]++
	}
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	for i := 0; i < month; i++ {
		sum += doy[i]
	}
	sum += day
	return sum
}
