package p10

func isMatch(s string, p string) bool {
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dp(s, 0, p, 0, memo)
}

func dp(s string, i int, p string, j int, memo [][]int) bool {
	m, n := len(s), len(p)
	if j == n {
		return i == m
	}
	if m == i {
		if (n-j)%2 == 1 {
			return false
		}
		for ; j < n-1; j = j + 2 {
			if p[j+1] != '*' {
				return false
			}
		}
		return true
	}
	if memo[i][j] != -1 {
		return memo[i][j] == 1
	}
	var res bool
	memo[i][j] = 0
	if s[i] == p[j] || p[j] == '.' {
		//匹配
		if j < n-1 && p[j+1] == '*' {
			res = dp(s, i, p, j+2, memo) || dp(s, i+1, p, j, memo)
		} else {
			res = dp(s, i+1, p, j+1, memo)
		}

	} else {
		if j < n-1 && p[j+1] == '*' {
			res = dp(s, i, p, j+2, memo)
		} else {
			res = false
		}
	}
	if res {
		memo[i][j] = 1
	}
	return res
}
