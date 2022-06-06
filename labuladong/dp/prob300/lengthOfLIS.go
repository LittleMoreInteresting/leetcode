package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0,1,0,3,2,3}))
	fmt.Println(lengthOfLIS([]int{7,7,7,7,7,7,7}))
}

func lengthOfLIS(nums []int) int {
	//base case
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	
	for i := 0; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
