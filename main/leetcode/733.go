package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if newColor == oldColor {
		return image
	}
	var stack [][]int
	stack = append(stack, []int{sr, sc})
	for len(stack) != 0 {
		tr := stack[0][0]
		tc := stack[0][1]
		image[tr][tc] = newColor
		stack = stack[1:]
		if tr >= 1 && image[tr-1][tc] == oldColor {
			image[tr-1][tc] = newColor
			stack = append(stack, []int{tr - 1, tc})
		}
		if tr < len(image)-1 && image[tr+1][tc] == oldColor {
			image[tr+1][tc] = newColor
			stack = append(stack, []int{tr + 1, tc})
		}
		if tc >= 1 && image[tr][tc-1] == oldColor {
			image[tr][tc-1] = newColor
			stack = append(stack, []int{tr, tc - 1})
		}
		if tc < len(image[0])-1 && image[tr][tc+1] == oldColor {
			image[tr][tc+1] = newColor
			stack = append(stack, []int{tr, tc + 1})
		}
	}
	return image
}
