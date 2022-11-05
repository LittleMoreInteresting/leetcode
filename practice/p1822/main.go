package main

import "fmt"

func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
}

func arraySign(nums []int) int {
	res := 1
	for i := range nums {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			res *= -1
		}
	}
	return res
}
