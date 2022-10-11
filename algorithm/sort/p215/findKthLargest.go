package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 5, 9, 3, 3}
	k := findKthLargest(nums, 4)
	fmt.Println(k, nums)
}

func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	k = len(nums) - 1 - k
	for l <= r {
		p := Partition(nums, l, r)
		fmt.Println(p, l, r)
		fmt.Println(nums)
		if p < k {
			// 第 k 大的元素在 nums[p+1..hi] 中
			l = p + 1
		} else if p > k {
			// 第 k 大的元素在 nums[lo..p-1] 中
			r = p - 1
		} else {
			// 找到第 k 大元素
			return nums[p]
		}
	}
	return -1
}

func Partition(nums []int, low, height int) int {
	i, pivot := low, nums[height]
	for j := low; j < height; j++ {
		if nums[j] < pivot {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
		}
	}
	nums[i], nums[height] = nums[height], nums[i]
	return i
}
