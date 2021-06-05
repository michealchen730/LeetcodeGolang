package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	mp := make(map[int]*Employee)
	for _, v := range employees {
		mp[v.Id] = v
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		imp := mp[i].Importance
		for _, v := range mp[i].Subordinates {
			imp += dfs(v)
		}
		return imp
	}
	return dfs(id)
}
