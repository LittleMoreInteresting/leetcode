package main

import "fmt"

func pivotIndex(nums []int) int {
	n := len(nums)
	prev := make([]int, n+1)
	prev[0] = 0
	for i := 1; i <= n; i++ {
		prev[i] = prev[i-1] + nums[i-1]
	}
	for i := 1; i < len(prev); i++ {
		left := prev[i-1] - prev[0]
		right := prev[n] - prev[i]
		if left == right {
			return i - 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
	nums = []int{1, 2, 3}
	fmt.Println(pivotIndex(nums))
	nums = []int{2, 1, -1}
	fmt.Println(pivotIndex(nums))
}