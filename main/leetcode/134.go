package main

//这题最开始的思路是枚举所有加油站
//这里应该有一个结论：
//如果A->B->C->D走不通了
//那么下一个起点应该是D以后的点

func canCompleteCircuit(gas []int, cost []int) int {

	for i := 0; i < len(gas); i++ {
		gas[i] -= cost[i]
	}

	var checkComplete func(start int) int

	checkComplete = func(start int) int {
		g := gas[start]
		destination := start
		for g >= 0 {
			start++
			start = start % len(gas)
			if start == destination {
				return -1
			}
			g += gas[start]
		}
		return start + 1
	}

	//枚举所有加油站
	//next是给出的下一个作为出发点的加油站
	for i := 0; i < len(gas); i++ {
		next := checkComplete(i)
		if next == -1 {
			return i
		} else if next <= i {
			//如果下一站是它本身或者序号上小于它，那么直接返回-1
			return -1
		} else {
			i = next - 1
		}
	}
	return -1
}
