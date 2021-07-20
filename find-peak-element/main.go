package main

import "fmt"

func main() {
	nums := []int{1, 2}
	peak := find_peak_element(nums)
	fmt.Println(peak)
}

func find_peak_element(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	peak := 0
	step := len(nums) - 1
	for i := 0; i < step; i++ {
		peak = i
		if nums[i] > nums[i+1] {
			return peak
		}
	}
	return peak + 1
}
