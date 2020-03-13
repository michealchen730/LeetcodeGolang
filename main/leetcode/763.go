package main

func partitionLabels(S string) []int {
	var res []int
	arr := make([][]int, 26)
	used := make([]int, 26)
	for i := 0; i < len(S); i++ {
		temp := S[i] - 'a'
		if len(arr[temp]) <= 1 {
			arr[temp] = append(arr[temp], i)
		} else {
			arr[temp][1] = i
		}
	}
	for i := 0; i < len(S); i++ {
		temp := S[i] - 'a'
		used[temp] = 1
		if len(arr[temp]) == 1 {
			res = append(res, 1)
		} else {
			span := []int{arr[temp][0], arr[temp][1]}
			for j := span[0]; j <= span[1]; j++ {
				temp2 := S[j] - 'a'
				if used[temp2] == 0 {
					used[temp2] = 1
					if len(arr[temp2]) > 1 && arr[temp2][1] > span[1] {
						span[1] = arr[S[j]-'a'][1]
					}
				}
			}
			res = append(res, span[1]-span[0]+1)
			i = span[1]
		}
	}
	return res
}
