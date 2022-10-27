package main

import (
	"fmt"

	"leetcode/labuladong/dp/public"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray00([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArrayWin([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray00(nums []int) int {
	l := len(nums)
	dp := make([]int, l) //定义：dp[i] 记录以 nums[i] 为结尾的「最大子数组和」
	dp[0] = nums[0]
	for i := 1; i < l; i++ {
		dp[i] = public.Max(nums[i], nums[i]+dp[i-1])
	}
	return public.Max(dp...)
}

func maxSubArray(nums []int) int {
	l := len(nums)
	dp_0 := nums[0]
	dp, res := 0, dp_0
	for i := 1; i < l; i++ {
		dp = public.Max(nums[i], nums[i]+dp_0)
		dp_0 = dp
		res = public.Max(res, dp)
	}
	return res
}

//滑动窗口
func maxSubArrayWin(nums []int) int {
	left, right := 0, 0
	winMax := 0
	res := nums[0]
	for right < len(nums) {
		winMax += nums[right]
		right++
		if winMax > res {
			res = winMax
		}
		for winMax < 0 {
			winMax -= nums[left]
			left++
		}
	}
	return res
}
