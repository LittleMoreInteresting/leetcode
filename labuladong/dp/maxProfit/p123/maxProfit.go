package p123

import "math"

func maxProfit(prices []int) int {
	n := len(prices)
	dp_i_20, dp_i_21, dp_i_10, dp_i_11 := 0, math.MinInt, 0, math.MinInt
	for i := 0; i < n; i++ {
		dp_i_20 = Max(dp_i_20, dp_i_21+prices[i])
		dp_i_21 = Max(dp_i_21, dp_i_10-prices[i])
		dp_i_10 = Max(dp_i_10, dp_i_11+prices[i])
		dp_i_11 = Max(dp_i_11, -prices[i])
	}
	return dp_i_20
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
