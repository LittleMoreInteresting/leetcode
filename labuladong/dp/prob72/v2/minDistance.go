package main

import (
	"fmt"
	"leetcode/labuladong/dp/public"
)

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention","execution"))
	fmt.Println(minDistance("intention","intention"))
}
//
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := InitDp(m+1,n+1) // 定义：s1[0..i] 和 s2[0..j] 的最小编辑距离是 dp[i+1][j+1]
	for i := 1; i < m; i++ {
		dp[i][0] = i
	}
	for j := 1;j<n;j++{
		dp[0][j] = j
	}
	for i := 1;i<=m;i++ {
		for j := 1;j<=n;j++ {
			if word1[i-1] == word2[j-1]{
				dp[i][j] = dp[i-1][j-1]
			}else{
				dp[i][j] = public.Min(
					dp[i-1][j-1]+1,
					dp[i-1][j]+1,
					dp[i][j-1]+1)
			}
		}
	}
	for _,v := range dp {
		fmt.Println(v)
	}
	return dp[m][n] 
}

// 初始化 m*n 切片
func InitDp(m,n int) [][]int {
	ch := make([]int, n*m)
	dp := make([][]int,m)
	for i := range dp {
		dp[i],ch = ch[:n],ch[n:]
	}
	return dp
}