package main

func removeDuplicates(nums []int) int {
	n := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}
