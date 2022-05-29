package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	// 计算前缀和
	prevSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		prevSum[i] = prevSum[i-1] + nums[i-1]
	}
	// 前缀和索引
	valIndex := make(map[int]int, len(prevSum))
	for i := 0; i < len(prevSum); i++ {
		v := prevSum[i] % k
		if _, ok := valIndex[v]; !ok {
			valIndex[v] = i
		}
	}
	for i := 1; i < len(prevSum); i++ {
		need := prevSum[i] % k
		if idx, ok := valIndex[need]; ok && (i-idx) >= 2 {
			return true
		}
	}
	return false
}

// 官方
func checkSubarraySum2(nums []int, k int) bool {
    m := len(nums)
    if m < 2 {
        return false
    }
    mp := map[int]int{0: -1}
    remainder := 0
    for i, num := range nums {
        remainder = (remainder + num) % k
        if prevIndex, has := mp[remainder]; has {
            if i-prevIndex >= 2 {
                return true
            }
        } else {
            mp[remainder] = i
        }
    }
    return false
}

func main() {
	nums := []int{0,0}
	fmt.Println(checkSubarraySum(nums, 1))
	fmt.Println(checkSubarraySum2(nums, 1))
}