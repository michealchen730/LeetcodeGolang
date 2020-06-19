package main

type SubrectangleQueries struct {
	mat [][]int
}

func Constructor1476(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{mat: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.mat[i][j] = newValue
		}
	}

}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.mat[row][col]
}
