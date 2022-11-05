package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}

func rob0(nums []int) int {
	l := len(nums)
	dp := make([]int, l+2)

	for i := l - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], dp[i+2]+nums[i])
	}
	return dp[0]
}
func rob(nums []int) int {
	l := len(nums)
	dp_1, dp_2 := 0, 0
	dp_max := 0
	for i := l - 1; i >= 0; i-- {
		dp_max = max(dp_1, dp_2+nums[i])
		dp_2 = dp_1
		dp_1 = dp_max
	}
	return dp_max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
