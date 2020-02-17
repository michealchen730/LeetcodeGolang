package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start==destination {
		return 0
	}
	length1,length2:=0,0
	if start>destination{
		start,destination=destination,start
	}
	for i:=0; i<len(distance); i++ {
		if i>=start&&i<destination{
			length1+=distance[i]
		}else {
			length2+=distance[i]
		}
	}
	if length1>length2 {
		return length2
	}else {
		return length1
	}
}
