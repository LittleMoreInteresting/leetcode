package main

import "fmt"

func main() {
	fmt.Println(champagneTower1(6, 3, 1))
}

func champagneTower1(poured int, query_row int, query_glass int) float64 {
	row := []float64{float64(poured)}
	for i := 1; i <= query_row; i++ {
		nextRow := make([]float64, i+1)
		for j, volume := range row {
			if volume > 1 {
				nextRow[j] += (volume - 1) / 2
				nextRow[j+1] += (volume - 1) / 2
			}
		}
		fmt.Println(nextRow)
		row = nextRow
	}
	if row[query_glass] < 1 {
		return row[query_glass]
	}
	return 1
}
