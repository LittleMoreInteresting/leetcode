package main

import (
	"leetcode/labuladong/dp/public"
)

func main() {

}
func minDistance(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([]int, n+1) // s1[0...i] 与 s2[0...j]
	for i := 1; i <= m; i++ {
		dp_i_j := 0
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = 1 + dp_i_j
			} else {
				dp[j] = public.Max(dp[j], dp[j-1])
			}
			dp_i_j = temp
		}
	}
	return m-dp[n]+n-dp[n]
}