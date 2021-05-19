package main

type ParkingSystem struct {
	nums [3]int
	used [3]int
}

func Constructor1603(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		nums: [3]int{big, medium, small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.nums[carType-1]-this.used[carType-1] > 0 {
		this.used[carType-1]++
		return true
	}
	return false
}
