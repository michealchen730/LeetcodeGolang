package main

type UndergroundSystem struct {
	mp  map[int]string
	mp2 map[string][]int
	mp3 map[int]int
}

//func Constructor() UndergroundSystem {
//	return UndergroundSystem{
//		mp:make(map[int]string),
//		mp2:make(map[string][]int),
//		mp3:make(map[int]int),
//	}
//}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.mp[id] = stationName
	this.mp3[id] = t
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	name := this.mp[id] + " " + stationName
	if len(this.mp2[name]) == 0 {
		this.mp2[name] = make([]int, 2)
	}
	this.mp2[name][0] += t - this.mp3[id]
	this.mp2[name][1]++
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	return float64(this.mp2[startStation+" "+endStation][0]) / float64(this.mp2[startStation+" "+endStation][1])
}
