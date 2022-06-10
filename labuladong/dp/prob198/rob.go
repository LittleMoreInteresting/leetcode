package main

import "fmt"

func main() {
	nums := []int{2, 1, 7, 9, 3, 1}
	fmt.Println(rob(nums))
	nums = []int{2,7,9,3,1}
	fmt.Println(rob(nums))
}
//Step 1 
//自顶向下的动态规划
func rob(nums []int) int {
	mem = make(map[int]int)
	return dp(nums, 0)
}
//Step 2
var mem  map[int]int // 递归备忘录
func dp(nums []int, start int) int {
	if v,ok := mem[start];ok {
		return v
	}
	if start >= len(nums) {
		return 0
	}
	res := max(dp(nums, start+1), nums[start]+dp(nums, start+2))
	mem[start] = res
	return res
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

