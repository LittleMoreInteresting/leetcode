package main

import (
	"fmt"
	"math"
)

// 暴力枚举
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount <= 0 {
		return -1
	}
	result := math.MaxInt64
	for _, coin := range coins {
		sub := coinChange(coins, amount-coin)
		if sub == -1 {
			continue
		}
		result = int(math.Min(float64(result), float64(sub+1)))
	}
	if result == math.MaxInt64 {
		result = -1
	}
	return result
}

func main() {
	coins := []int{1, 2, 5}
	ret := coinChange(coins, 11)
	fmt.Println(ret)
}
