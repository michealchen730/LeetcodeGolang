package main

func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1)==0||len(list2)==0{
		return nil
	}
	minsum:=-1
	dic:=make(map[string]int)
	minress:=make(map[int][]string)
	var longl,shortl []string
	if len(list1)>len(list2) {
		longl,shortl=list1,list2
	}else {
		longl,shortl=list2,list1
	}
	for k,v:=range longl{
		dic[v]=k
	}
	for k,v:=range shortl{
		if t,ok:=dic[v];ok {
			if minsum==-1 {
				minsum=k+t
				minress[minsum]=append(minress[minsum],v)
			}else{
				if k+t<=minsum {
					minsum=k+t
					minress[minsum]=append(minress[minsum],v)
				}
			}
		}
	}
	return minress[minsum]
}