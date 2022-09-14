package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	r, l := 0, int(math.Sqrt(float64(c)))
	for r <= l {
		res := r*r + l*l
		if res == c {
			return true
		}
		if res > c {
			l--
		} else {
			r++
		}
	}
	fmt.Println(r, ":", l)
	return false
}

func main() {
	fmt.Println(judgeSquareSum(2))
}
