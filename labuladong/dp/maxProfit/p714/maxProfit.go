package p714

import (
	"math"
)

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt
	for i := 0; i < n; i++ {
		temp := dp_i_0
		dp_i_0 = Max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = Max(dp_i_1, temp-prices[i]-fee)
	}
	return dp_i_0
}

func Max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
