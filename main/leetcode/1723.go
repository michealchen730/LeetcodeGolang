package main

import (
	"sort"
)

func minimumTimeRequired(jobs []int, k int) int {
	res := 0
	for _, v := range jobs {
		res += v
	}
	// tmp表示当前是jobs的第几个
	var dfs1723 func(workers []int, tmp, limit int)
	dfs1723 = func(workers []int, tmp, limit int) {
		//tmp<0说明成功过关
		if tmp < 0 {
			//找到当前的最大值
			ans := 0
			for _, v := range workers {
				ans = max(ans, v)
			}
			//与res进行比较
			if ans < res {
				res = ans
			}
			return
		}
		//这里不是等于，limit应该是成功组合中算到的最大值
		if res <= limit {
			return
		}
		for k, v := range workers {
			//剪枝,只要当前工人和他的前一位工作量一致，就可以直接跳过
			if k != 0 && v == workers[k-1] {
				continue
			}
			//只要符合条件就加入组合
			if v+jobs[tmp] <= limit {
				workers[k] += jobs[tmp]
				dfs1723(workers, tmp-1, limit)
				workers[k] -= jobs[tmp]
			}
		}
	}
	sort.Ints(jobs)
	low := jobs[len(jobs)-1]
	high := res
	for low <= high {
		mid := low + (high-low)/2
		workers := make([]int, k)
		dfs1723(workers, len(jobs)-1, mid)
		//这里可以进一步优化
		if res <= mid {
			high = res - 1
		} else {
			low = mid + 1
		}
	}
	return res
}
