package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumDeleteSum("sea","eat"));
	fmt.Println(mem)
}

var mem [][]int // 备忘录

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	mem = InitDp(m, n)
	return dp(s1,0,s2,0)
}

func dp(s1 string, i int, s2 string, j int) int {
	
	res := 0
	if i == len(s1) {
		for ;j<len(s2);j++ {
			res += int(s2[j])
		}
		return res
	}
	if j == len(s2) {
		for ;i<len(s1);i++ {
			res += int(s1[i])
		}
		return res
	}

	if mem[i][j] != -1 {
		return mem[i][j]
	}

	if s1[i] == s2[j] {
		mem[i][j] = dp(s1, i+1, s2, j+1)
	} else {
		mem[i][j] = Min(
			int(s1[i]) + dp(s1, i+1, s2, j),
			int(s2[j]) + dp(s1, i, s2, j+1),
		)
	}
	return mem[i][j]
}

// 初始化 m*n 切片
func InitDp(m, n int) [][]int {
	ch := make([]int, n*m)
	for idx := range ch {
		ch[idx] = -1
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i], ch = ch[:n], ch[n:]
	}
	return dp
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