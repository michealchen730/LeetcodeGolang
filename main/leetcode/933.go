package main

type RecentCounter struct {
	pngs []int
}

func Constructor933() RecentCounter {
	return RecentCounter{pngs: []int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.pngs = append(this.pngs, t)
	last := max(t-3000, 0)
	newStart := 0
	for i := 0; i < len(this.pngs); i++ {
		if this.pngs[i] >= last {
			newStart = i
			break
		}
	}
	this.pngs = this.pngs[newStart:]
	return len(this.pngs)
}
