package main

import (
	"fmt"
	"math"
)

//增加备忘录
var mem map[int]int

func coinChange(coins []int, amount int) int {
	mem = make(map[int]int)
	return dep(coins, amount)
}

func dep(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if v, ok := mem[amount]; ok {
		return v
	}
	result := math.MaxInt64
	for _, coin := range coins {
		sub := dep(coins, amount-coin)
		if sub == -1 {
			continue
		}
		result = int(math.Min(float64(result), float64(sub+1)))
	}
	if result == math.MaxInt64 {
		result = -1
	}
	mem[amount] = result
	return result
}

func main() {
	coins := []int{1, 2, 5}
	ret := coinChange(coins, 11)
	fmt.Println(ret)
	fmt.Println(mem)
}
