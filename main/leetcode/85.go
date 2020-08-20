package main

//这里看了第四种解法
//思想其实是这样的，一旦有最大矩形，那么这个矩形内必定存在某一个点A
//点A位于矩形的底边上，往上扩展的长度为h，往左右两边扩展的宽度为w
//此时h*w==s

//如果我们只从单个点出发，那么先往上扩展再往左右扩展得到的h‘和w’，有可能并不是这个点所能构成的最大矩形
//因为动态规划考虑所有情况，因此第四种解法就有了必然成立性
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	//没有使用官方题解的三数组（那个写法是真的省空间复杂度）
	r, c := len(matrix), len(matrix[0])
	dp := make([][]dpNodes, r)
	for k, _ := range dp {
		dp[k] = make([]dpNodes, c)
	}

	res := 0
	for i := 0; i < r; i++ {
		cur_left, cur_right := 0, c
		//因为是逐列遍历，所以cur_left,cur_right会递增和递减
		for j := 0; j < c; j++ {
			//当matrix[i][j]==1时才会受上一行影响
			if matrix[i][j] == '1' {
				dp[i][j].height = 1
				if i > 0 {
					dp[i][j].height += dp[i-1][j].height
					dp[i][j].left = max(dp[i-1][j].left, cur_left)
				} else {
					dp[i][j].left = max(0, cur_left)
				}
			} else {
				dp[i][j].left = 0
				cur_left = j + 1
			}
		}
		for j := c - 1; j >= 0; j-- {
			//当matrix[i][j]==1时才会受上一行影响
			if matrix[i][j] == '1' {
				if i > 0 {
					dp[i][j].right = min(dp[i-1][j].right, cur_right)
				} else {
					dp[i][j].right = min(c, cur_right)
				}
			} else {
				dp[i][j].right = cur_right
				cur_right = j
			}

			res = max(res, (dp[i][j].right-dp[i][j].left)*dp[i][j].height)
		}
	}
	return res
}

type dpNodes struct {
	height int
	left   int
	right  int
}
