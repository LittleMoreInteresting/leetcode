package main

import (
	"math"
)

func main() {

}

func bestCoordinate(towers [][]int, radius int) []int {
	var xMax, yMax, pointX, pointY, maxQuality int
	for _, t := range towers {
		xMax = max(xMax, t[0])
		yMax = max(yMax, t[1])
	}
	for x := 0; x <= xMax; x++ {
		for y := 0; y <= yMax; y++ {
			quality := 0
			for _, t := range towers {
				d := pointDis(x, y, t[0], t[1])
				if d <= float64(radius) {
					quality += int(float64(t[2]) / (1 + d))
				}
			}
			if quality > maxQuality {
				pointX, pointY, maxQuality = x, y, quality
			}
		}
	}
	return []int{pointX, pointY}
}

func pointDis(x1, y1, x2, y2 int) float64 {
	dis := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	return math.Sqrt(float64(dis))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
