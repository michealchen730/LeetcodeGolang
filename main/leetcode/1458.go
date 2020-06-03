package main

func maxDotProduct(nums1 []int, nums2 []int) int {
	mat := make([][]int, len(nums1))
	for k, _ := range mat {
		mat[k] = make([]int, len(nums2))
	}
	mat[0][0] = nums1[0] * nums2[0]
	t1, t2 := mat[0][0], mat[0][0]
	for i := 0; i < len(nums1); i++ {
		t1 = max(t1, nums1[i]*nums2[0])
		mat[i][0] = t1
	}
	for j := 0; j < len(nums2); j++ {
		t2 = max(t2, nums1[0]*nums2[j])
		mat[0][j] = t2
	}
	res := max(t1, t2)
	for i := 1; i < len(nums1); i++ {
		for j := 1; j < len(nums2); j++ {
			mat[i][j] = nums1[i] * nums2[j]
			mat[i][j] = max(mat[i-1][j], mat[i][j])
			mat[i][j] = max(mat[i][j-1], mat[i][j])
			mat[i][j] = max(mat[i][j], mat[i-1][j-1])
			mat[i][j] = max(mat[i][j], nums1[i]*nums2[j]+mat[i-1][j-1])
			res = max(mat[i][j], res)
		}
	}
	return res
}
