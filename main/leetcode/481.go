package main

//这里有一点，使用make([]int,4,n+3)，后面衔接append操作，
//会明显快于make([]int,n+3)，然后对数组的每一个元素进行赋值

func magicalString(n int) int {
	if n == 0 {
		return 0
	}
	arr := make([]int, 4, n+3)
	arr[1], arr[2], arr[3] = 1, 2, 2
	temp, idx, val := 4, 3, 1
	cnt := 1
	for temp <= n {
		for i := 0; i < arr[idx]; i++ {
			if val == 1 {
				if temp <= n {
					cnt++
				}
				arr = append(arr, 1)
			} else {
				arr = append(arr, 2)
			}

			temp++
		}
		if val == 1 {
			val = 2
		} else {
			val = 1
		}
		idx++
	}
	return cnt
}

//func magicalString(n int) int {
//	if n==0{
//		return 0
//	}
//	arr:=make([]int,100010)
//	arr[1],arr[2],arr[3]=1,2,2
//	temp,idx,val:=4,3,1
//	cnt:=1
//	for temp<=n{
//		for i:=0;i<arr[idx];i++{
//			if temp<=n&&val==1{
//				cnt++
//			}
//			arr[temp]=val
//			temp++
//		}
//		if val==1{
//			val=2
//		}else{
//			val=1
//		}
//		idx++
//	}
//	return cnt
//}
