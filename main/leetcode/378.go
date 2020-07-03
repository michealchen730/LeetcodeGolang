package main

func kthSmallest378(matrix [][]int, k int) int {
	//题解给的二分查找非常巧妙
	left, right := matrix[0][0], matrix[len(matrix)-1][len(matrix)-1]
	//这里其实是需要证明的
	//我们看到这里会最终让算法执行到left==right为止，为什么此时的left一定是矩阵中的元素呢，这里需要证明一下
	for left < right {
		mid := left + (right-left)/2
		if checkNums378(matrix, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	//可以参考这个例子：[[1,5,9],[5,11,13],[12,13,15]]，k=2和k=3
	return left
}

func checkNums378(matrix [][]int, k, x int) bool {
	i, j := 0, matrix[len(matrix)][0]
	nums := 0
	for i < len(matrix) && j >= 0 {
		if matrix[j][i] <= x {
			nums += j + 1 //j代表的是当前所在的行下标，因为行也是按照升序排列，所以如果matrix[j][i]<=x，那么必然所在的那一列都小于等于x
			i++
		} else {
			j--
		}
	}
	return nums >= k
}
