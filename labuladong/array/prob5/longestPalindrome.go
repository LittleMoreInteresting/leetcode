package prob5

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

// 动态规划解法

func longestPalindrome_DP(s string) string {
	//1 定义DP数组  dp[i][j] 代表子串 s[i]~s[j] 是否为回文字符串
	length := len(s)
	if length == 0 {
		return ""
	}
	dp := initDp(length)
	//2 定义初始状态
	start, longest := 0, 1
	dp[length-1][length-1] = true
	for i := 0; i < length-1; i++ {
		dp[i][i] = true

		dp[i][i+1] = s[i] == s[i+1]
		if dp[i][i+1] {
			start = i
			longest = 2
		}
	}

	//3 状态转移方程 dp[i][j] = dp[i+1][j-1] & s[i] == s[j]
	for i := length - 1; i >= 0; i-- {
		for j := i + 2; j < length; j++ {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			if dp[i][j] && j-i+1 > longest {
				start = i
				longest = j - i + 1
			}
		}
	}
	return s[start : start+longest]
}

func initDp(len int) [][]bool {
	dp := make([][]bool, len)
	dpTemp := make([]bool, len*len)
	for i := 0; i < len; i++ {
		dp[i], dpTemp = dpTemp[0:len], dpTemp[len:]
	}
	return dp
}
