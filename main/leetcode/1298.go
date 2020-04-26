package main

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	var bcbp []int //表示能被打开的盒子
	bcnbp := make(map[int]int)
	haskey := make(map[int]int) //表示当前持有的钥匙
	res := 0                    //当前获得的糖果总数
	//初始化
	for _, v := range initialBoxes {
		if status[v] == 1 {
			bcbp = append(bcbp, v)
		} else {
			bcnbp[v] = 1
		}
	}
	for len(bcbp) != 0 {
		length := len(bcbp)
		for i := 0; i < length; i++ {
			//拿到当前盒子
			tempBox := bcbp[i]
			//拿到当前盒子的糖果
			res += candies[tempBox]
			//拿到当前盒子里的钥匙
			for _, v := range keys[tempBox] {
				haskey[v] = 1
			}
			//拿到当前盒子里的盒子
			for _, v := range containedBoxes[tempBox] {
				//盒子可以打开，放进bcbp中，如果需要钥匙打开，就放进bcnbp中
				if status[v] == 1 {
					bcbp = append(bcbp, v)
				} else {
					bcnbp[v] = 1
				}
			}
		}
		//一轮结束以后，打开所有能打开的盒子
		for k := range haskey {
			if bcnbp[k] != 0 {
				bcbp = append(bcbp, k)
				delete(haskey, k)
				delete(bcnbp, k)
			}
		}
		bcbp = bcbp[length:]
	}
	return res
}
