package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition0([]int{1, 5, 11, 5}))
}

func canPartition0(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	n := len(nums)
	sum = sum / 2
	dp := InitDp(n+1, sum+1)
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	for _, i2 := range dp {
		fmt.Println(i2)
	}
	return dp[n][sum]
}
func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	n := len(nums)
	sum = sum / 2
	dp := make([]bool, sum+1)
	dp[0] = true

	for i := 0; i < n; i++ {
		for j := sum; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}
	for _, i2 := range dp {
		fmt.Println(i2)
	}
	return dp[sum]
}

// 初始化 m*n 切片
func InitDp(m, n int) [][]bool {
	ch := make([]bool, n*m)
	dp := make([][]bool, m)
	for i := range dp {
		dp[i], ch = ch[:n], ch[n:]
	}
	return dp
}
