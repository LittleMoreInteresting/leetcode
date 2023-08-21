package main

import (
	"fmt"
	"math"
)

func main() {
	var n, s int
	_, _ = fmt.Scan(&n, &s)

	minP := math.MaxInt32 - s
	var total int64

	for i := 0; i < n; i++ {
		var p, y int
		_, _ = fmt.Scan(&p, &y)

		minP = int(math.Min(float64(minP+s), float64(p)))
		total += int64(minP * y)
	}
	fmt.Println(total)
}
