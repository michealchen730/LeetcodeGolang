package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	res1:=make([]int,numCourses)
	for i:=0;i<numCourses;i++{
		res1[i]=i
	}
	if len(prerequisites)==0 {
		return res1
	}

	var res2 []int
	var stack []int
	inDeg:=make(map[int]int)
	next:=make(map[int][]int)
	for i:=0;i<len(prerequisites);i++{
		inDeg[prerequisites[i][0]]++
		if _,ok:=inDeg[prerequisites[i][1]]; !ok{
			inDeg[prerequisites[i][1]]=0
		}
		next[prerequisites[i][1]]=append(next[prerequisites[i][1]], prerequisites[i][0])
	}

	for i:=0;i<numCourses;i++{
		if _,ok:=inDeg[i]; !ok{
			res2= append(res2, i)
		}
	}

	for k,v:=range inDeg{
		if v==0 {
			stack= append(stack, k)
		}
	}

	for len(stack)!=0{
		temp:=stack[0]
		stack=stack[1:]
		res2=append(res2,temp)
		for _,v:=range next[temp]{
			if inDeg[v]==1{
				stack= append(stack, v)
			}
			inDeg[v]--
		}
	}

	if numCourses==len(res2) {
		return res2
	}else{
		return []int{}
	}
}