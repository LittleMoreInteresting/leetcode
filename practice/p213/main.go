package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
}

func rob(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	return max(robLine(nums, 1, l-1),
		robLine(nums, 0, l-2),
	)
}

func robLine(nums []int, start, end int) int {
	dp_1, dp_2 := 0, 0
	dp_max := 0
	for i := end; i >= start; i-- {
		dp_max = max(dp_1, dp_2+nums[i])
		dp_2 = dp_1
		dp_1 = dp_max
	}
	return dp_max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
