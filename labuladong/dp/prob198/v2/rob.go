package main

import "fmt"

func main() {
	nums := []int{2, 1, 7, 9, 3, 1}
	fmt.Println(rob(nums))
	nums = []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}

//自底向上
func Rob0(nums []int) int {
	// dp[i] = x 表示：
    // 从第 i 间房子开始抢劫，最多能抢到的钱为 x
    // base case: dp[n] = 0
	l := len(nums)
	dp := make([]int,l+2)
	for i := l-1;i>=0;i-- {
		dp[i] = max(nums[i]+dp[i+2],dp[i+1])
	}
	return dp[0]
}
// 进一步简化
func rob(nums []int) int {
	// dp[i] = x 表示：
    // 从第 i 间房子开始抢劫，最多能抢到的钱为 x
    // base case: dp[n] = 0
	l := len(nums)
	dp_i,dp_i_1,dp_i_2 := 0,0,0
	for i := l-1;i>=0;i-- {
		dp_i = max(nums[i]+dp_i_2,dp_i_1)
		//状态转义
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
