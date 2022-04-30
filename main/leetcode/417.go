package main

func pacificAtlantic(heights [][]int) [][]int {
	nextStep := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var pacific, atlantic [][]int

	for i := 0; i < len(heights); i++ {
		pacific = append(pacific, []int{i, 0})
		atlantic = append(atlantic, []int{i, len(heights[0]) - 1})
	}

	for i := 0; i < len(heights[0]); i++ {
		pacific = append(pacific, []int{0, i})
		atlantic = append(atlantic, []int{len(heights) - 1, i})
	}

	p := make(map[int]bool)
	a := make(map[int]bool)

	for len(pacific) != 0 {
		length := len(pacific)
		for i := 0; i < length; i++ {
			tmp := pacific[i]
			p[tmp[0]*len(heights[0])+tmp[1]] = true
			for _, v := range nextStep {
				row, col := tmp[0]+v[0], tmp[1]+v[1]
				if row >= 0 && row < len(heights) && col >= 0 && col < len(heights[0]) && !p[row*len(heights[0])+col] && heights[row][col] >= heights[tmp[0]][tmp[1]] {
					pacific = append(pacific, []int{row, col})
				}
			}
		}
		pacific = pacific[length:]
	}

	for len(atlantic) != 0 {
		length := len(atlantic)
		for i := 0; i < length; i++ {
			tmp := atlantic[i]
			a[tmp[0]*len(heights[0])+tmp[1]] = true
			for _, v := range nextStep {
				row, col := tmp[0]+v[0], tmp[1]+v[1]
				if row >= 0 && row < len(heights) && col >= 0 && col < len(heights[0]) && !a[row*len(heights[0])+col] && heights[row][col] >= heights[tmp[0]][tmp[1]] {
					atlantic = append(atlantic, []int{row, col})
				}
			}
		}
		atlantic = atlantic[length:]
	}
	var res [][]int

	for k, _ := range p {
		if a[k] {
			res = append(res, []int{k / len(heights[0]), k % len(heights[0])})
		}
	}
	return res
}
