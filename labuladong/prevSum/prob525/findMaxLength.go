package main

import "fmt"

func findMaxLength(nums []int) int {
	prevSum := make([]int, len(nums)+1)

	for i := 1; i < len(prevSum); i++ {
		var sum int = 1
		if nums[i-1] == 0 {
			sum = -1
		}
		prevSum[i] = prevSum[i-1] + sum
	}
	var valIndex = make(map[int]int, len(prevSum))
	for i := 0; i < len(prevSum); i++ {
		if _, ok := valIndex[prevSum[i]]; !ok {
			valIndex[prevSum[i]] = i
		}
	}
	var res int
	for i := 1; i < len(prevSum); i++ {
		need := prevSum[i]
		if val, ok := valIndex[need]; ok {
			res = max(res, i-val)
		}
	}
	fmt.Println(prevSum)
	fmt.Println(valIndex)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{0, 1,0 }
	fmt.Println(findMaxLength(nums))
}