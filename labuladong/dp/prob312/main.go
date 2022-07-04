package main

import "fmt"

func main() {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
}
func maxCoins(nums []int) int {
	n := len(nums)
	points := make([]int, n+2)
	points[0], points[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		points[i] = nums[i-1]
	}
	dp := InitDp(n+2, n+2)
	for i := n; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = Max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
			}
		}
	}
	return dp[0][n+1]
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

func Max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
