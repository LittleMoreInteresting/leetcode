package main

import "fmt"

/**
给定一个数组 nums ，将其划分为两个连续子数组 left 和 right， 使得：

left 中的每个元素都小于或等于 right 中的每个元素。
left 和 right 都是非空的。
left 的长度要尽可能小。
在完成这样的分组后返回 left 的 长度 。

用例可以保证存在这样的划分方法。
*/
func main() {
	nums := []int{5, 0, 3, 8, 6}
	fmt.Println(partitionDisjoint(nums))
	fmt.Println(partitionDisjoint1(nums))
	nums = []int{1, 1, 1, 0, 6, 12}
	fmt.Println(partitionDisjoint(nums))
	fmt.Println(partitionDisjoint1(nums))
}

func partitionDisjoint1(nums []int) int {
	length := len(nums)
	if length == 2 {
		return 1
	}
	lo, loMax, cur := 0, nums[0], nums[0]
	for i := 1; i < length-1; i++ {
		if nums[i] > cur {
			cur = nums[i]
		}
		if nums[i] < loMax {
			loMax = cur
			lo = i
		}
	}
	return lo + 1
}

func partitionDisjoint(nums []int) int {
	n := len(nums)
	leftMax, leftPos, curMax := nums[0], 0, nums[0]
	for i := 1; i < n-1; i++ {
		curMax = max(curMax, nums[i])
		if nums[i] < leftMax {
			leftMax = curMax
			leftPos = i
		}
	}
	return leftPos + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
