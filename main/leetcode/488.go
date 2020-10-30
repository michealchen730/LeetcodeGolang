package main

//此题参考了题解区的思路
//通过暴力的思路解题会导致超时的结果

//剪枝的主要思路为（来自题解区）：
//遍历时，如果后面的球与前面的球颜色不一样，在这里尝试插入一个后面颜色的球
//如果相邻的两个球颜色相同，考虑在中间插入一个其他颜色的球，将他们分割

func findMinStep(board string, hand string) int {
	//mp用来存储我们手上的球
	mp := make(map[byte]int)
	for k, _ := range hand {
		mp[hand[k]]++
	}
	res := -1

	var dfs func(b []byte)

	dfs = func(b []byte) {

		//检查mp的状态，如果我们没球了，那么直接返回
		cnt := 0
		for _, v := range mp {
			cnt += v
		}
		//如果祖玛消耗完了，那也返回
		if len(b) == 0 {

			if res == -1 {
				res = len(hand) - cnt
			} else {
				res = min(res, len(hand)-cnt)
			}

			return
		}

		if cnt == 0 {
			return
		}

		//开始dfs
		//选择位置插球
		flag := true

		for i := 0; i < len(b); i++ {
			//如果当前元素与它前一个元素颜色不同，那么试图插入一个前一元素的颜色(需要我们手中有这一颜色的球)
			if i > 0 && b[i] != b[i-1] && mp[b[i-1]] > 0 {
				flag = false
				//手中同颜色的球数目减去一
				mp[b[i-1]]--

				//预存下b数组
				pre := make([]byte, len(b))
				copy(pre, b)

				//如果存在消除，那么就消除
				//b中添加一个元素
				pre = append(pre[:i], append([]byte{pre[i-1]}, pre[i:]...)...)
				handleZumaArr(&pre)

				dfs(pre)

				mp[b[i-1]]++
			}
			//如果已经到达末尾
			if i == len(b)-1 && mp[b[i]] > 0 {
				flag = false

				mp[b[i]]--

				//预存下b数组
				pre := make([]byte, len(b))
				copy(pre, b)

				//如果存在消除，那么就消除
				//b中添加一个元素
				pre = append(pre, pre[i])
				handleZumaArr(&pre)

				//if len(pre)==0&&len(b)==4{
				//	fmt.Println(b)
				//}

				dfs(pre)

				mp[b[i]]++
			}

			//如果前后两个元素相同，那么考虑往中间添加一个不同元素的球
			if i > 0 && b[i] == b[i-1] {
				for k, v := range mp {
					if k != b[i] && v > 0 {
						flag = false

						mp[k]--

						//预存下b数组
						pre := make([]byte, len(b))
						copy(pre, b)

						//b中添加一个元素
						pre = append(pre[:i], append([]byte{k}, pre[i:]...)...)
						dfs(pre)

						mp[k]++
					}
				}
			}
		}
		//不存在添加球的可能性，也直接返回
		if flag {
			return
		}
	}

	bytes := []byte(board)
	dfs(bytes)

	return res
}

func handleZumaArr(b *[]byte) {
	for i := 0; i < len(*b); i++ {
		j := i
		for j < len(*b) {
			if (*b)[j] == (*b)[i] {
				j++
			} else {
				break
			}
		}
		if j-i >= 3 {
			*b = append((*b)[:i], (*b)[j:]...)
			i = -1
		}

	}

}
