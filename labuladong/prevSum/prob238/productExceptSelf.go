package main

import "fmt"

func productExceptSelf(nums []int) []int {
	l := len(nums)
	prefix := make([]int, l)
	prefix[0] = nums[0]
	suffix := make([]int, l)
	suffix[l-1] = nums[l-1]
	for i := 1; i < l; i++ {
		prefix[i] = prefix[i-1] * nums[i]
		suffix[l-1-i] = suffix[l-i] * nums[l-1-i]
	}

	res := make([]int, l)
	res[0], res[l-1] = suffix[1], prefix[l-2]
	for i := 1; i < l-1; i++ {
		res[i] = prefix[i-1] * suffix[i+1]
	}
	return res
}

func main() {

	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
	nums = []int{-1,1,0,-3,3}
	fmt.Println(productExceptSelf(nums))
}