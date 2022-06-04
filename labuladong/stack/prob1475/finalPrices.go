package main

import "fmt"

func finalPrices(prices []int) []int {
	n := len(prices)
	res := make([]int, n)
	nextPrice := nextLessOrEqualElement(prices)
	for i := 0; i < n; i++ {
		if nextPrice[i] == -1 {
			res[i] = prices[i]
		} else {
			res[i] = prices[i] - nextPrice[i]
		}
	}
	return res
}

func nextLessOrEqualElement(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	st := make([]int, 0)
	for i := l - 1; i >= 0; i-- {

		for len(st) > 0 && st[0] > nums[i] {
			st = Pop(st)
		}
		if len(st) == 0 {
			res[i] = -1
		} else {
			res[i] = st[0]
		}
		st = Push(st, nums[i])
	}
	return res
}
func Pop(nums []int) []int {
	return nums[1:]
}

func Push(nums []int, num int) []int {
	temp := []int{num}
	temp = append(temp, nums...)
	return temp
}
func main() {
	fmt.Println(finalPrices([]int{8,4,6,2,3}))
}