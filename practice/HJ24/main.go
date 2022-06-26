package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	nums := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&nums[i])
	}
	left, right := make([]int, count), make([]int, count)
	left[0], right[count-1] = 1, 1
	for i := 0; i < count; i++ {
		left[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				left[i] = Max(left[i], left[j]+1)
			}
		}
	}
	for i := count - 1; i >= 0; i-- {
		right[i] = 1
		for j := count - 1; j > i; j-- {
			if nums[i] > nums[j] {
				right[i] = Max(right[i], right[j]+1)
			}
		}
	}
	max := 1
	for i := 0; i < count; i++ {
		res := left[i] + right[i] - 1
		if max < res {
			max = res
		}
	}
	fmt.Print(count - max)
}
func Max(a ...int) int {
	max := a[0]
	for _, v := range a[1:] {
		if v > max {
			max = v
		}
	}
	return max
}
