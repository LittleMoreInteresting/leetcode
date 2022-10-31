package main

import "fmt"

/**
给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。

由于答案可能很大，因此 返回答案模 10^9 + 7 。
*/
func main() {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
}

func sumSubarrayMins(arr []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	monoStack := []int{}
	for i, x := range arr {
		for len(monoStack) > 0 && x <= arr[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = i + 1
		} else {
			left[i] = i - monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	monoStack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(monoStack) > 0 && arr[i] < arr[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = n - i
		} else {
			right[i] = monoStack[len(monoStack)-1] - i
		}
		monoStack = append(monoStack, i)
	}
	fmt.Println(left)
	fmt.Println(right)
	for i, x := range arr {
		ans = (ans + left[i]*right[i]*x) % mod
	}
	return
}

func sumSubarrayMinsDP(arr []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(arr)
	monoStack := []int{}
	dp := make([]int, n)
	for i, x := range arr {
		for len(monoStack) > 0 && arr[monoStack[len(monoStack)-1]] > x {
			monoStack = monoStack[:len(monoStack)-1]
		}
		k := i + 1
		if len(monoStack) > 0 {
			k = i - monoStack[len(monoStack)-1]
		}
		dp[i] = k * x
		if len(monoStack) > 0 {
			dp[i] += dp[i-k]
		}
		ans = (ans + dp[i]) % mod
		monoStack = append(monoStack, i)
	}
	fmt.Println(dp)
	return
}

func sumSubarrayMinsW(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += backtrack(arr, i)
	}
	return sum
}

func backtrack(nums []int, win int) int {
	sum := 0
	if win == 0 {
		for n := range nums {
			sum += nums[n]
		}
		return sum
	}
	left, right := 0, 0
	for right < len(nums) {
		right++
		if right-left > win {
			fmt.Println(right, "-", left, "=", nums[left:right])
			sum += min(nums[left:right])
			left++
		}
	}
	return sum
}

func min(a []int) int {
	m := a[0]
	for _, v := range a[1:] {
		if v < m {
			m = v
		}
	}
	return m
}
