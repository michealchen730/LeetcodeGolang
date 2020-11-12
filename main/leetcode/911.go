package main

type TopVotedCandidate struct {
	winner []int
	times  []int
}

func Constructor911(persons []int, times []int) TopVotedCandidate {
	winner := make([]int, len(times))
	mx := 0
	mp := make(map[int]int)
	for k, v := range persons {
		mp[v]++
		if mp[v] >= mx {
			winner[k] = v
			mx = mp[v]
		} else {
			winner[k] = winner[k-1]
		}
	}
	return TopVotedCandidate{
		winner: winner,
		times:  times,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	low, high := 0, len(this.times)-1
	if t >= this.times[high] {
		return this.winner[high]
	}
	for low < high {
		mid := low + (high-low)/2
		if t < this.times[mid] {
			high = mid
		} else {
			if t < this.times[mid+1] {
				return this.winner[mid]
			}
			low = mid
		}
	}
	return -1
}
