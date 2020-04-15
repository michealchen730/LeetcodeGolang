package main

func commonChars(A []string) []string {
	arr1 := make([]int, 26)
	for k, w := range A {
		arr2 := make([]int, 26)
		for _, b := range w {
			arr2[b-'a']++
		}
		if k == 0 {
			copy(arr1, arr2)
		} else {
			for k, _ := range arr1 {
				arr1[k] = min(arr1[k], arr2[k])
			}
		}
	}
	var res []string
	for k, v := range arr1 {
		for i := 0; i < v; i++ {
			res = append(res, string('a'+k))
		}
	}
	return res
}
