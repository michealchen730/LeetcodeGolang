package main

func findWords(board [][]byte, words []string) []string {
	//根据题目中的测试用例，最终返回的结果中是不可以出现相同的单词的
	mp := make(map[string]bool)

	var res []string
	t := Trie{}
	for _, v := range words {
		t.Insert(v)
	}
	isSearched := make([][]bool, len(board))
	for k, _ := range isSearched {
		isSearched[k] = make([]bool, len(board[0]))
	}
	//使用深度优先去实现搜索单词
	var dfs func(i, j int, path *[]byte)
	dfs = func(i, j int, path *[]byte) {
		//终止条件1，超出边界
		if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 {
			return
		}

		//终止条件2，回到之前的位置
		if isSearched[i][j] {
			return
		}

		//前两个限制均通过，那么
		*path = append(*path, board[i][j])
		isSearched[i][j] = true

		//终止条件3，没有当前前缀
		prefix := string(*path)
		if !t.StartsWith(prefix) {
			*path = (*path)[:len(*path)-1]
			isSearched[i][j] = false
			return
		}

		//否则继续递归
		if t.Search(prefix) && !mp[prefix] {
			mp[prefix] = true
			res = append(res, prefix)
		}
		dfs(i+1, j, path)
		dfs(i-1, j, path)
		dfs(i, j+1, path)
		dfs(i, j-1, path)
		*path = (*path)[:len(*path)-1]
		isSearched[i][j] = false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, &[]byte{})
		}
	}
	return res
}
