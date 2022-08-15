package p313

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	m := len(primes)
	pointers := make([]int, m)
	nums := make([]int, m)
	for i := range nums {
		nums[i] = 1
	}
	for i := 1; i <= n; i++ {
		minNum := math.MaxInt64
		for j := range pointers {
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j := range nums {
			if nums[j] == minNum {
				pointers[j]++
				nums[j] = dp[pointers[j]] * primes[j]
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
