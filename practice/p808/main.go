package p808

func soupServings(n int) float64 {
	n = (n + 24) / 25
	if n >= 179 {
		return 1
	}
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, n+1)
	}
	var dfs func(int, int) float64
	dfs = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			return 0.5
		}
		if a <= 0 {
			return 1
		}
		if b <= 0 {
			return 0
		}
		dv := &dp[a][b]
		if *dv > 0 {
			return *dv
		}
		res := (dfs(a-4, b) + dfs(a-3, b-1) +
			dfs(a-2, b-2) + dfs(a-1, b-3)) / 4
		*dv = res
		return res
	}
	return dfs(n, n)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/*作者：力扣官方题解
链接：https://leetcode.cn/problems/soup-servings/solutions/1981704/fen-tang-by-leetcode-solution-0yxs/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
