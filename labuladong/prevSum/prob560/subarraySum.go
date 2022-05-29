package main

import "fmt"

func subarraySum(nums []int, k int) int {
	prevSum := make([]int, len(nums)+1)
	prevSum[0] = 0
	count := make(map[int]int, len(nums)+1)
	count[0] = 1
	result := 0
	for i := 1; i < len(prevSum); i++ {
		prevSum[i] = prevSum[i-1] + nums[i-1]
		need := prevSum[i] - k
		if v, ok := count[need]; ok {
			result += v
		}
		if v, ok := count[prevSum[i]]; ok {
			count[prevSum[i]] = v + 1
		} else {
			count[prevSum[i]] = 1
		}
	}

	return result
}

//官方
func subarraySum1(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

func main() {
	nums := []int{1,2,3}
	fmt.Println(subarraySum(nums, 3))
	fmt.Println(subarraySum1(nums, 3))
}