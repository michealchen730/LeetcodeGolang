package main

func minCostToMoveChips(chips []int) int {
	numeven:=0
	numodd:=0
	for i:=0;i<len(chips);i++{
		if chips[i]%2==0 {
			numeven++
		}else{
			numodd++
		}
	}
	if numeven>numodd {
		return numodd
	}else {
		return numeven
	}
}
