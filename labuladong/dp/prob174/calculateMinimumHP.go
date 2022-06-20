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
}

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
