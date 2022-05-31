package main

import "fmt"

func splitArray(nums []int, m int) int {
	left := 0
	right := 0
	for _, v := range nums {
		left = max(left, v)
		right += v
	}
	for left <= right {
		mid := left + (right-left)/2
		if needDays(nums, mid) <= m {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func needDays(weights []int, x int) int {
	days := 0
	sum := 0
	for _, v := range weights {
		if (sum + v) > x {
			days++
			sum = v
		} else {
			sum += v
		}
	}
	if sum > 0 {
		days++
	}
	return days
}
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(splitArray(weights, 5)) //15
	weights = []int{3, 2, 2, 4, 1, 4}
	fmt.Println(splitArray(weights, 3)) //6
	weights = []int{1, 2, 3, 1, 1}
	fmt.Println(splitArray(weights, 4)) //3
}