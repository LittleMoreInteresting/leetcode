package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 2, 4, -1, 9, -2, 4}
	quicksort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println(search(nums, -1))
}

func quicksort(nums []int, lo, he int) {
	if lo >= he {
		return
	}
	p := partition(nums, lo, he)
	quicksort(nums, lo, p-1)
	quicksort(nums, p+1, he)
}

func partition(nums []int, lo, he int) int {
	point := nums[lo]
	i, j := lo+1, he
	for i <= j {
		for i < he && nums[i] <= point {
			i++
		}
		for j > lo && nums[j] > point {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[j], nums[lo] = nums[lo], nums[j]
	return j
}

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			hi = mid - 1
			continue
		}
		if nums[mid] < target {
			lo = mid + 1
		}
	}
	return -1
}
