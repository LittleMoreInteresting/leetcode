package main

import (
	"fmt"
)

func nextGreaterElements503(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	stack := make([]int, 0)
	for i := 2*l - 1; i >= 0; i-- {
		idx := i % l
		for len(stack) > 0 && stack[0] <= nums[idx] {
			stack = pop(stack)
		}
		if len(stack) == 0 {
			res[idx] = -1
		}else {
			res[idx] = stack[0]
		}
		stack = push(stack, nums[idx])
	}
	return res
}

func nextGreaterElements(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	stack := make([]int, 0)
	for i := l - 1; i >= 0; i-- {
		
		for len(stack) > 0 && stack[0] <= nums[i] {
			stack = pop(stack)
		}
		if len(stack) == 0 {
			res[i] = -1
		}else {
			res[i] = stack[0]
		}
		stack = push(stack, nums[i])
	}
	return res
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greater := nextGreaterElements(nums2)
	gIdx := make(map[int]int, len(nums1))
	for i, v := range greater {
		gIdx[nums2[i]] = v;
	}
	res := make([]int,len(nums1))
	for i:=0;i<len(nums1);i++ {
		res[i] = gIdx[nums1[i]]
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
	//fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	//fmt.Println(nextGreaterElements([]int{1,2,3,4,3}))
	fmt.Println(nextGreaterElements503([]int{100,1,11,1,120,111,123,1,-1,-100}))
	fmt.Println(nextGreaterElement([]int{4,1,2},[]int{1,3,4,2}))
	fmt.Println(nextGreaterElement([]int{2,4},[]int{1,2,3,4}))

}