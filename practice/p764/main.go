package p764

func main() {

}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	dp := InitDp(n)
	banned := map[int]bool{}
	for _, p := range mines {
		banned[p[0]*n+p[1]] = true
	}
	for i := 0; i < n; i++ {
		count := 0
		/* left */
		for j := 0; j < n; j++ {
			if banned[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}
		count = 0
		/* right */
		for j := n - 1; j >= 0; j-- {
			if banned[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		count := 0
		/* up */
		for j := 0; j < n; j++ {
			if banned[j*n+i] {
				count = 0
			} else {
				count++
			}
			dp[j][i] = min(dp[j][i], count)
		}
		count = 0
		/* down */
		for j := n - 1; j >= 0; j-- {
			if banned[j*n+i] {
				count = 0
			} else {
				count++
			}
			dp[j][i] = min(dp[j][i], count)
			res = max(res, dp[j][i])
		}
	}
	return res
}

// 初始化 m*n 切片
func InitDp(n int) [][]int {
	ch := make([]int, n*n)
	dp := make([][]int, n)
	for i := range dp {
		dp[i], ch = ch[:n], ch[n:]
		for j := range dp[i] {
			dp[i][j] = n
		}
	}

	return dp
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func orderOfLargestPlusSign1(n int, mines [][]int) (ans int) {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = n
		}
	}
	banned := map[int]bool{}
	for _, p := range mines {
		banned[p[0]*n+p[1]] = true
	}
	for i := 0; i < n; i++ {
		count := 0
		/* left */
		for j := 0; j < n; j++ {
			if banned[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}
		count = 0
		/* right */
		for j := n - 1; j >= 0; j-- {
			if banned[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}
	}
	for i := 0; i < n; i++ {
		count := 0
		/* up */
		for j := 0; j < n; j++ {
			if banned[j*n+i] {
				count = 0
			} else {
				count++
			}
			dp[j][i] = min(dp[j][i], count)
		}
		count = 0
		/* down */
		for j := n - 1; j >= 0; j-- {
			if banned[j*n+i] {
				count = 0
			} else {
				count++
			}
			dp[j][i] = min(dp[j][i], count)
			ans = max(ans, dp[j][i])
		}
	}
	return ans
}

//作者：力扣官方题解
