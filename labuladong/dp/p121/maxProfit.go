package p121

import (
	"math"

	"leetcode/labuladong/dp/public"
)

func maxProfit(prices []int) int {
	n := len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt
	for i := 0; i < n; i++ {
		dp_i_0 = public.Max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = public.Max(dp_i_1, -prices[i])
	}
	return dp_i_0
}
