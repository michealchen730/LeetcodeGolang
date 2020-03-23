package main

//type MaxQueue struct {
//	queue []int
//	max   []int
//}
//
//func Constructor() MaxQueue {
//	return MaxQueue{}
//}
//
//func (this *MaxQueue) Max_value() int {
//	if len(this.queue) == 0 {
//		return -1
//	}
//	return this.max[0]
//}
//
//func (this *MaxQueue) Push_back(value int) {
//	this.queue = append(this.queue, value)
//	this.max = append(this.max, value)
//	i := 0
//	for i < len(this.max) {
//		if this.max[i] < value {
//			this.max = append(this.max[0:i], this.max[i+1:]...)
//			i--
//		}
//		i++
//	}
//}
//func (this *MaxQueue) Pop_front() int {
//	if len(this.queue) == 0 {
//		return -1
//	}
//	temp := this.queue[0]
//	this.queue = this.queue[1:]
//	if temp == this.max[0] {
//		this.max = this.max[1:]
//	}
//	return temp
//}
