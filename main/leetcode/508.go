package main

func findFrequentTreeSum(root *TreeNode) []int {
	mp := make(map[int]int)
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		a, b := dfs(root.Left), dfs(root.Right)
		sum := a + b + root.Val
		mp[sum]++
		return sum
	}

	dfs(root)

	var res []int
	cnt := 0
	for k, v := range mp {
		if v > cnt {
			cnt = v
			res = []int{k}
		} else if v == cnt {
			res = append(res, k)
		}
	}
	return res
}
