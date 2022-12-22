package p813

func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	prefix := make([]float64, n+1)
	for i, x := range nums {
		prefix[i+1] = prefix[i] + float64(x)
	}
	dp := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = prefix[i] / float64(i)
	}
	for j := 2; j <= k; j++ {
		for i := n; i >= j; i-- {
			for x := j - 1; x < i; x++ {
				dp[i] = max(dp[i], dp[x]+(prefix[i]-prefix[x])/float64(i-x))
			}
		}
	}
	return dp[n]
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

/*作者：力扣官方题解
链接：https://leetcode.cn/problems/largest-sum-of-averages/solutions/1993132/zui-da-ping-jun-zhi-he-de-fen-zu-by-leet-09xt/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
