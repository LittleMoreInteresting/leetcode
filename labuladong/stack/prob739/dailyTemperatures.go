package main

import "fmt"

func dailyTemperatures(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	stack := make([]int, 0)
	for i := l - 1; i >= 0; i-- {
		
		for len(stack) > 0 && nums[stack[0]] <= nums[i] {
			stack = pop(stack)
		}
		if len(stack) == 0 {
			res[i] = 0
		}else {
			res[i] = stack[0] - i
		}
		
		stack = push(stack, i)
	}
	return res
}
func pop(nums []int) []int {
	return nums[1:]
}

func push(nums []int, num int) []int {
	temp := []int{num}
	temp = append(temp, nums...)
	return temp
}
func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30,60,90}))
}