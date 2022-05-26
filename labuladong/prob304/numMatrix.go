package main

type NumMatrix struct {
	prevSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {

	row, column := len(matrix), len(matrix[0])
	prevSum := make([][]int, row+1)
	for i := range prevSum {
		prevSum[i] = make([]int, column+1)
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			prevSum[i][j] = prevSum[i-1][j] + prevSum[i][j-1] + matrix[i-1][j-1] - prevSum[i-1][j-1]
		}
	}

	return NumMatrix{prevSum: prevSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.prevSum[row2+1][col2+1] - this.prevSum[row1][col2+1] - this.prevSum[row2+1][col1] + this.prevSum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */