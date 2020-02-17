package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses==0||len(prerequisites)==0 {
		return true
	}
	//出入度，两个map，一个用来存后继结点，一个用来存点的入度
	inDeg:=make(map[int]int)
	neNodes:=make(map[int][]int)
	var stack []int
	for i:=0;i<len(prerequisites);i++{
		inDeg[prerequisites[i][0]]++
		if _,ok:=inDeg[prerequisites[i][1]]; !ok {
			inDeg[prerequisites[i][1]]=0
		}
		neNodes[prerequisites[i][1]]=append(neNodes[prerequisites[i][1]], prerequisites[i][0])
	}

	for k,v:=range inDeg {
		if v==0 {
			stack= append(stack, k)
		}
	}

	for len(stack)!=0{
		for _,v:=range neNodes[stack[0]]{
			if inDeg[v]==1 {
				stack= append(stack, v)
			}
			inDeg[v]--
		}
		delete(inDeg,stack[0])
		stack=stack[1:]
	}

	if len(inDeg)==0{
		return true
	}

	//if res==numCourses {
	//	return true
	//}
	return false
}
