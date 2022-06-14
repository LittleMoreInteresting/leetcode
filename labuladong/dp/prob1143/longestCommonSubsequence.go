package main

import (
	"fmt"
	"leetcode/labuladong/dp/public"
)

func main() {

	fmt.Println(longestCommonSubsequence("mhunuzqrkzsnidwbun","szulspmhwpazoxijwbq"))
	fmt.Println(longestCommonSubsequence("abcde","ace"))
	fmt.Println(longestCommonSubsequence2("abcde","ace"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := InitDp(m+1, n+1) //dp[i+1][j+1] s1[0...i] 与 s2[0...j] 的 最大LCS
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = public.Max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}
func longestCommonSubsequence2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([]int,n+1)// s1[0...i] 与 s2[0...j]
	//1 dp[j] == dp[i-1][j]
	//2 dp[j-1] == dp[i][j-1]  用 dp_i_j 存 dp[i-1][j-1]
	for i := 1; i <= m; i++ {
		dp_i_j := 0
		for j := 1; j <= n; j++ {
			temp := dp[j];
			if text1[i-1] == text2[j-1] {
				dp[j] = 1 + dp_i_j
			} else {
				dp[j] = public.Max(dp[j], dp[j-1])
			}
			dp_i_j = temp
		}
		fmt.Println(dp)
	}
	
	return dp[n]
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