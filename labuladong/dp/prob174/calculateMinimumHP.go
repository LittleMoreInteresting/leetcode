package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(calculateMinimumHP([][]int{{0, -3}}))
	fmt.Println(calculateMinimumHP([][]int{{0}}))
	fmt.Println(calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}))
	fmt.Println(calculateMinimumHP([][]int{{-3, 5}}))

	fmt.Println(calculateMinimumHPV2([][]int{{0, -3}}))
	fmt.Println(calculateMinimumHPV2([][]int{{0}}))
	fmt.Println(calculateMinimumHPV2([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}))
	fmt.Println(calculateMinimumHPV2([][]int{{-3, 5}}))
}

func calculateMinimumHPV2(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := InitDp(m+1, n+1)
	if dungeon[m-1][n-1] >= 0 {
		dp[m][n] = 1
	} else {
		dp[m][n] = -dungeon[m-1][n-1] + 1
	}
	for i := m - 1; i > 0; i-- {
		if dungeon[i-1][n-1] >= dp[i+1][n] {
			dp[i][n] = 1
		} else {
			dp[i][n] = dp[i+1][n] - dungeon[i-1][n-1]
		}
	}
	for j := n - 1; j > 0; j-- {
		if dungeon[m-1][j-1] >= dp[m][j+1] {
			dp[m][j] = 1
		} else {
			dp[m][j] = dp[m][j+1] - dungeon[m-1][j-1]
		}
	}
	for i := m - 1; i > 0; i-- {
		for j := n - 1; j > 0; j-- {
			var left, down int
			if dungeon[i-1][j-1] >= dp[i][j+1] {
				left = 1
			} else {
				left = dp[i][j+1] - dungeon[i-1][j-1]
			}
			if dungeon[i-1][j-1] >= dp[i+1][j] {
				down = 1
			} else {
				down = dp[i+1][j] - dungeon[i-1][j-1]
			}
			dp[i][j] = Min(left, down)
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[1][1]
}
func InitDp(m, n int) [][]int {
	ch := make([]int, n*m)
	dp := make([][]int, m)
	for i := range dp {
		dp[i], ch = ch[:n], ch[n:]
	}
	return dp
}

///*****************************************************
var meme map[string]int

func calculateMinimumHP(dungeon [][]int) int {
	meme = make(map[string]int)
	return dp(dungeon, 0, 0)
}

func dp(dungeon [][]int, i, j int) int {
	m, n := len(dungeon), len(dungeon[0])
	if i == m-1 && j == n-1 {
		if dungeon[i][j] >= 0 {
			return 1
		} else {
			return -dungeon[i][j] + 1
		}
	}
	if i == m || j == n {
		return math.MaxInt
	}
	key := strconv.Itoa(i) + "-" + strconv.Itoa(j)
	if v, ok := meme[key]; ok {
		return v
	}

	res := Min(dp(dungeon, i, j+1), dp(dungeon, i+1, j)) - dungeon[i][j]
	if res <= 0 {
		meme[key] = 1
	} else {
		meme[key] = res
	}
	return meme[key]
}

func Min(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < m {
			m = nums[i]
		}
	}
	return m
}
