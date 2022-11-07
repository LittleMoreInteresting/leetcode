package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 4, -1, 9, -2, 4}
	quicksort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println(search(nums, -1))
}

func quicksort(nums []int, l, h int) {
	if l >= h {
		return
	}
	p := partition(nums, l, h)
	quicksort(nums, l, p-1)
	quicksort(nums, p+1, h)
}

func partition(nums []int, l, h int) int {
	point := nums[l]
	i, j := l+1, h
	for i <= j {
		for i < h && nums[i] <= point {
			i++
		}
		for j > l && nums[j] > point {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[l], nums[j] = nums[j], nums[l]
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
