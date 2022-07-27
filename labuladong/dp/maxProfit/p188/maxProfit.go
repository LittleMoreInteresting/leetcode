package p188

import (
	"math"
)

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if k > n/2 {
		return maxProfitAny(prices)
	}
	dp := initDp(n, k+1)
	for i := 0; i < n; i++ {
		for j := k; j >= 1; j-- {
			if i-1 == -1 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
			} else {
				dp[i][j][0] = Max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
				dp[i][j][1] = Max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			}
		}
	}

	return dp[n-1][k][0]
}

func initDp(n, k int) [][][]int {
	dp := make([][][]int, n)
	for i, _ := range dp {
		dp[i] = make([][]int, k)
		for j, _ := range dp[i] {
			if j == 0 {
				dp[i][j] = []int{0, math.MinInt}
			} else {
				dp[i][j] = []int{0, 0}
			}
		}
	}
	return dp
}

func maxProfitAny(prices []int) int {
	n := len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt
	for i := 0; i < n; i++ {
		temp := dp_i_0
		dp_i_0 = Max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = Max(dp_i_1, temp-prices[i])
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
