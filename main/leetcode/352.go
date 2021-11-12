package main

type SummaryRanges struct {
	Intervals [][]int
}

//func Constructor() SummaryRanges {
//	intervals:=[][]int{{-2,-2},{10002,10002}}
//	return SummaryRanges{
//		Intervals: intervals,
//	}
//}

func (this *SummaryRanges) AddNum(val int) {
	n := len(this.Intervals)
	start, end := 0, n-1
	for start < end {
		mid := (start + end + 1) >> 1
		if this.Intervals[mid][0] <= val {
			// 命中了也要后移
			start = mid
		} else {
			// 没有命中立马前移
			end = mid - 1
		}
	}
	pre := this.Intervals[end]
	next := this.Intervals[end+1]
	if pre[0] <= val && val <= pre[1] || next[0] <= val && val <= next[1] {
	} else if pre[1]+1 == val && val == next[0]-1 {
		pre[1] = next[1]
		this.Intervals = append(this.Intervals[:end+1], this.Intervals[end+2:]...)
	} else if pre[1]+1 == val {
		pre[1] = val
	} else if next[0]-1 == val {
		next[0] = val
	} else {
		this.Intervals = append(this.Intervals[:end+1], append([][]int{{val, val}}, this.Intervals[end+1:]...)...)
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.Intervals[1 : len(this.Intervals)-1]
}
