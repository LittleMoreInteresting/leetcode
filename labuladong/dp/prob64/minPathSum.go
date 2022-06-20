package main

import (
	"fmt"
	"math"
	"strconv"
)

// https://leetcode.cn/problems/minimum-path-sum/

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum1([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

//自底向上的迭代解法 + 空间压缩
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] += dp[j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = Min(dp[j], dp[j-1]) + grid[i][j]
		}
	}
	return dp[n-1]
}

// 自底向上的迭代解法
func minPathSum1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := InitDp(m, n)
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] += dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] += dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	for i := range dp {
		fmt.Println(dp[i])
	}
	return dp[m-1][n-1]
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

// 自顶向下动态规划=++++++++++++++++++++++++++++++++++++++++++++++++++++++
var meme map[string]int

func minPathSum0(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	meme = make(map[string]int)
	return dp(grid, m-1, n-1)
}

func dp(grid [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	if i < 0 || j < 0 {
		return math.MaxInt
	}
	key := strconv.Itoa(i) + "-" + strconv.Itoa(j)
	if v, ok := meme[key]; ok {
		return v
	}
	meme[key] = Min(
		dp(grid, i, j-1),
		dp(grid, i-1, j),
	) + grid[i][j]
	return meme[key]
}

func Min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
