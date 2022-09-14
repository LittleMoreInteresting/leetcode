package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	l, r := 0, x
	sq := -1
	for l <= r {
		m := l + (r-l)/2
		res := m * m
		if res == x {
			return m
		}
		if res > x {
			r = m - 1
		} else {
			sq = m
			l = m + 1
		}
	}
	return sq
}
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}
func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(10))
	fmt.Println(mySqrt(11))
	fmt.Println(mySqrt(12))
}
