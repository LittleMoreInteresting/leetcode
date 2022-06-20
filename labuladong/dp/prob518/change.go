package main

import "fmt"

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}

func change(amount int, coins []int) int {
	n := len(coins)
	dp := InitDp(n+1, amount+1) //dp[N][amount]
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	for _, i2 := range dp {
		fmt.Println(i2)
	}
	return dp[n][amount]
}

// 初始化 m*n 切片
func InitDp(m, n int) [][]int {
	ch := make([]int, n*m)
	dp := make([][]int, m)
	for i := range dp {
		dp[i], ch = ch[:n], ch[n:]
	}
	return dp
}
