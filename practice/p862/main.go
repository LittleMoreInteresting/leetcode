package main

import (
	"fmt"
)

func main() {
	//fmt.Println(shortestSubarray([]int{1}, 1))
	//fmt.Println(shortestSubarray([]int{1, 2}, 4))
	fmt.Println(shortestSubarray([]int{2, -1, 2}, 3))
	fmt.Println(shortestSubarray([]int{56, -21, 56, 35, -9}, 61))
}
func shortestSubarray(nums []int, k int) int {
	length := len(nums)
	preSum := make([]int, length+1)
	preSum[0] = 0
	for i := 1; i <= length; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	minLen := length + 1
	q := []int{}
	for i, curSum := range preSum {
		for len(q) > 0 && curSum-preSum[q[0]] >= k {
			minLen = Min(minLen, i-q[0])
			q = q[1:]
		}
		for len(q) > 0 && preSum[q[len(q)-1]] >= curSum {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	if minLen == length+1 {
		return -1
	}
	return minLen
}

func Min(nums ...int) int {
	minVal := nums[0]
	for idx := range nums {
		if nums[idx] < minVal {
			minVal = nums[idx]
		}
	}
	return minVal
}
